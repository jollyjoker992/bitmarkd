// SPDX-License-Identifier: ISC
// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/bitmark-inc/bitmarkd/account"
	"github.com/bitmark-inc/bitmarkd/block"
	"github.com/bitmark-inc/bitmarkd/blockheader"
	"github.com/bitmark-inc/bitmarkd/fault"
	"github.com/bitmark-inc/bitmarkd/util"
	"github.com/bitmark-inc/bitmarkd/zmqutil"
	"github.com/bitmark-inc/exitwithstatus"
	"github.com/bitmark-inc/logger"
)

const (
	peerPublicKeyFilename  = "peer.public"
	peerPrivateKeyFilename = "peer.private"

	rpcCertificateKeyFilename = "rpc.crt"
	rpcPrivateKeyFilename     = "rpc.key"

	proofPublicKeyFilename      = "proof.public"
	proofPrivateKeyFilename     = "proof.private"
	proofLiveSigningKeyFilename = "proof.live"
	proofTestSigningKeyFilename = "proof.test"
)

// setup command handler
// commands that run to create key and certificate files
// these commands cannot access any internal database or states
func processSetupCommand(arguments []string) bool {

	command := "help"
	if len(arguments) > 0 {
		command = arguments[0]
		arguments = arguments[1:]
	}

	var arg map[string]string
	var err error
	if len(arguments) > 0 {
		arg, err = parseArgs(arguments)
		if nil != err {
			fmt.Printf("could not read arguments. error: %s\n", err)
			exitwithstatus.Exit(1)
		}
	}

	switch command {
	case "gen-peer-identity", "peer":
		publicKeyFilename := getFilenameWithDirectory(arg, peerPublicKeyFilename)
		privateKeyFilename := getFilenameWithDirectory(arg, peerPrivateKeyFilename)
		err := zmqutil.MakeKeyPair(publicKeyFilename, privateKeyFilename)
		if nil != err {
			fmt.Printf("generate private key: %q and public key: %q error: %s\n", privateKeyFilename, publicKeyFilename, err)
			exitwithstatus.Exit(1)
		}
		fmt.Printf("generated private key: %q and public key: %q\n", privateKeyFilename, publicKeyFilename)

	case "gen-rpc-cert", "rpc":
		certificateFilename := getFilenameWithDirectory(arg, rpcCertificateKeyFilename)
		privateKeyFilename := getFilenameWithDirectory(arg, rpcPrivateKeyFilename)

		addresses := []string{}
		addrArgs := arg["a"]
		if "" != addrArgs {
			addresses = strings.Split(addrArgs, ",")
		}

		err := makeSelfSignedCertificate("rpc", certificateFilename, privateKeyFilename, 0 != len(addresses), addresses)
		if nil != err {
			fmt.Printf("generate RPC key: %q and certificate: %q error: %s\n", privateKeyFilename, certificateFilename, err)
			exitwithstatus.Exit(1)
		}
		fmt.Printf("generated RPC key: %q and certificate: %q\n", privateKeyFilename, certificateFilename)

	case "gen-proof-identity", "proof":
		publicKeyFilename := getFilenameWithDirectory(arg, proofPublicKeyFilename)
		privateKeyFilename := getFilenameWithDirectory(arg, proofPrivateKeyFilename)
		err := zmqutil.MakeKeyPair(publicKeyFilename, privateKeyFilename)
		if nil != err {
			fmt.Printf("generate private key: %q and public key: %q error: %s\n", privateKeyFilename, publicKeyFilename, err)
			exitwithstatus.Exit(1)
		}

		liveSigningKeyFilename := getFilenameWithDirectory(arg, proofLiveSigningKeyFilename)
		testSigningKeyFilename := getFilenameWithDirectory(arg, proofTestSigningKeyFilename)
		version := "v2" // default is v2 twelve seed
		if v, ok := arg["v"]; ok {
			version = v
		}

		if err := makeSigningKey(version, false, liveSigningKeyFilename); nil != err {
			fmt.Printf("generate the signing key for livenet: %q error: %s\n", liveSigningKeyFilename, err)
			goto singingKeyFailed
		}
		if err := makeSigningKey(version, true, testSigningKeyFilename); nil != err {
			fmt.Printf(" generate the signing key for testnet: %q error: %s\n", testSigningKeyFilename, err)
			goto singingKeyFailed
		}

		fmt.Printf("generated private key: %q and public key: %q\n", privateKeyFilename, publicKeyFilename)
		fmt.Printf("generated signing keys: %q and %q\n", liveSigningKeyFilename, testSigningKeyFilename)
		goto done

	singingKeyFailed:
		_ = os.Remove(publicKeyFilename)
		_ = os.Remove(privateKeyFilename)
		exitwithstatus.Exit(1)

	case "dns-txt", "txt":
		return false // defer processing until configuration is read

	case "start", "run":
		return false // continue processing

		// case "block-times":
		// 	return false // defer processing until database is loaded

	case "block", "b", "save-blocks", "save", "load-blocks", "load", "delete-down", "dd":
		return false // defer processing until database is loaded

	default:
		switch command {
		case "help", "h", "?":
		case "", " ":
			fmt.Printf("error: missing command\n")
		default:
			fmt.Printf("error: no such command: %q\n", command)
		}

		fmt.Printf("supported commands:\n\n")
		fmt.Printf("  help                       		(h)      		- display this message\n\n")

		fmt.Printf("  gen-peer-identity d=[DIR]    		(peer)   		- create private key in: %q\n", "DIR/"+peerPrivateKeyFilename)
		fmt.Printf("                                        			and the public key in: %q\n", "DIR/"+peerPublicKeyFilename)
		fmt.Printf("\n")

		fmt.Printf("  gen-rpc-cert d=[DIR]         		(rpc)    		- create private key in:  %q\n", "DIR/"+rpcPrivateKeyFilename)
		fmt.Printf("    	                                    			and the certificate in: %q\n", "DIR/"+rpcCertificateKeyFilename)
		fmt.Printf("\n")

		fmt.Printf("  gen-rpc-cert d=[DIR] a=[IPs...]  			       	- create private key in:  %q\n", "DIR/"+rpcPrivateKeyFilename)
		fmt.Printf("                                        			and the certificate in: %q\n", "DIR/"+rpcCertificateKeyFilename)
		fmt.Printf("\n")

		fmt.Printf("  gen-proof-identity d=[DIR] v=[VER]   	(proof)  		- create private key in: %q\n", "DIR/"+proofPrivateKeyFilename)
		fmt.Printf("                                        			the public key in:     %q\n", "DIR/"+proofPublicKeyFilename)
		fmt.Printf("                                        			and signing keys in:  %q and: %q\n", "DIR/"+proofLiveSigningKeyFilename, "DIR/"+proofTestSigningKeyFilename)
		fmt.Printf("\n")

		fmt.Printf("  dns-txt                    		(txt)    		- display the data to put in a dbs TXT record\n")
		fmt.Printf("\n")

		fmt.Printf("  start                      		(run)    		- just run the program, same as no arguments\n")
		fmt.Printf("                                        			for convienience when passing script arguments\n")
		fmt.Printf("\n")

		fmt.Printf("  block S [E [FILE]]         		(b)      		- dump block(s) as a JSON structures to stdout/file\n")
		fmt.Printf("\n")

		fmt.Printf("  save-blocks FILE           		(save)   		- dump all blocks to a file\n")
		fmt.Printf("\n")

		fmt.Printf("  load-blocks FILE           		(load)   		- restore all blocks from a file\n")
		fmt.Printf("                                        			only runs if database is deleted first\n")
		fmt.Printf("\n")

		fmt.Printf("  delete-down NUMBER         		(dd)     		- delete blocks in descending order\n")
		fmt.Printf("\n")

		exitwithstatus.Exit(1)
	}

done:
	// indicate processing complete and prefor normal exit from main
	return true
}

func parseArgs(arguments []string) (map[string]string, error) {
	arg := make(map[string]string)

	for _, a := range arguments {
		i := strings.Index(a, "=")
		if -1 == i {
			return nil, fmt.Errorf("invalid arguments. it must be: key=value")
		}

		k := a[0:i]
		v := a[i+1:]
		if "" == k || "" == v {
			return nil, fmt.Errorf("invalid arguments. the key/value is empty")
		}
		arg[k] = v
	}

	return arg, nil
}

// configuration file enquiry commands
// have configuration file read and decoded, but nothing else
func processConfigCommand(arguments []string, options *Configuration) bool {

	command := "help"
	if len(arguments) > 0 {
		command = arguments[0]
	}

	switch command {
	case "dns-txt", "txt":
		dnsTXT(options)
	default: // unknown commands fall through to data command
		return false
	}

	// indicate processing complete and perform normal exit from main
	return true
}

// data command handler
// the internal block and transaction pools are enabled so these commands can
// access and/or change these databases
func processDataCommand(log *logger.L, arguments []string, options *Configuration) bool {

	command := "help"
	if len(arguments) > 0 {
		command = arguments[0]
		arguments = arguments[1:]
	}

	switch command {

	case "start", "run":
		return false // continue processing

	case "block", "b":
		if len(arguments) < 1 {
			exitwithstatus.Message("missing block number argument")
		}

		n, err := strconv.ParseUint(arguments[0], 10, 64)
		if nil != err {
			exitwithstatus.Message("error in block number: %s", err)
		}
		if n < 2 {
			exitwithstatus.Message("error: invalid block number: %d must be greater than 1", n)
		}

		output := "-"

		// optional end range
		nEnd := n
		if len(arguments) > 1 {

			nEnd, err = strconv.ParseUint(arguments[1], 10, 64)
			if nil != err {
				exitwithstatus.Message("error in ending block number: %s", err)
			}
			if nEnd < n {
				exitwithstatus.Message("error: invalid ending block number: %d must be greater than 1", n)
			}
		}

		if len(arguments) > 2 {
			output = strings.TrimSpace(arguments[2])
		}
		fd := os.Stdout

		if output != "" && output != "-" {
			fd, err = os.Create(output)
			if nil != err {
				exitwithstatus.Message("error: creating: %q error: %s", output, err)
			}
		}

		fmt.Fprintf(fd, "[\n")
		for ; n <= nEnd; n += 1 {
			block, err := dumpBlock(n)
			if nil != err {
				exitwithstatus.Message("dump block error: %s", err)
			}
			s, err := json.MarshalIndent(block, "  ", "  ")
			if nil != err {
				exitwithstatus.Message("dump block JSON error: %s", err)
			}

			fmt.Fprintf(fd, "  %s,\n", s)
		}
		fmt.Fprintf(fd, "{}]\n")
		fd.Close()

	case "save-blocks", "save":
		if len(arguments) < 1 {
			exitwithstatus.Message("missing file name argument")
		}
		filename := arguments[0]
		if "" == filename {
			exitwithstatus.Message("missing file name")
		}
		err := saveBinaryBlocks(filename)
		if nil != err {
			exitwithstatus.Message("failed writing: %q  error: %s", filename, err)
		}

	case "load-blocks", "load":
		if len(arguments) < 1 {
			exitwithstatus.Message("missing file name argument")
		}
		filename := arguments[0]
		if "" == filename {
			exitwithstatus.Message("missing file name")
		}
		err := restoreBinaryBlocks(filename)
		if nil != err {
			exitwithstatus.Message("failed writing: %q  error: %s", filename, err)
		}

	case "delete-down", "dd":
		// delete blocks down to a given block number
		if len(arguments) < 1 {
			exitwithstatus.Message("missing block number argument")
		}

		n, err := strconv.ParseUint(arguments[0], 10, 64)
		if nil != err {
			exitwithstatus.Message("error in block number: %s", err)
		}
		if n < 2 {
			exitwithstatus.Message("error: invalid block number: %d must be greater than 1", n)
		}
		err = block.DeleteDownToBlock(n)
		if nil != err {
			exitwithstatus.Message("block delete error: %s", err)
		}
		fmt.Printf("reduced height to: %d\n", blockheader.Height())

	default:
		exitwithstatus.Message("error: no such command: %s", command)

	}

	// indicate processing complete and perform normal exit from main
	return true
}

// print out the DNS TXT record
func dnsTXT(options *Configuration) {
	//   <TAG> a=<IPv4;IPv6> c=<PEER-PORT> r=<RPC-PORT> f=<SHA3-256(cert)> p=<PUBLIC-KEY>
	const txtRecord = `TXT "bitmark=v3 a=%s c=%d r=%d f=%x p=%x"` + "\n"

	rpc := options.ClientRPC

	keypair, err := tls.X509KeyPair([]byte(rpc.Certificate), []byte(rpc.PrivateKey))
	if nil != err {
		exitwithstatus.Message("error: cannot decode certificate: %q  error: %s", rpc.Certificate, err)
	}

	fingerprint := CertificateFingerprint(keypair.Certificate[0])

	if 0 == len(rpc.Announce) {
		exitwithstatus.Message("error: no rpc announce fields given")
	}

	rpcIP4, rpcIP6, rpcPort := getFirstConnections(rpc.Announce)
	if 0 == rpcPort {
		exitwithstatus.Message("error: cannot determine rpc port")
	}

	peering := options.Peering

	publicKey, err := zmqutil.ReadPublicKey(peering.PublicKey)
	if nil != err {
		exitwithstatus.Message("error: cannot read public key: %q  error: %s", peering.PublicKey, err)
	}

	if 0 == len(peering.Announce) {
		exitwithstatus.Message("error: no rpc announce fields given")
	}

	listenIP4, listenIP6, listenPort := getFirstConnections(peering.Announce)
	if 0 == listenPort {
		exitwithstatus.Message("error: cannot determine listen port")
	}

	IPs := ""
	if "" != rpcIP4 && rpcIP4 == listenIP4 {
		IPs = rpcIP4
	}
	if "" != rpcIP6 && rpcIP6 == listenIP6 {
		if "" == IPs {
			IPs = rpcIP6
		} else {
			IPs += ";" + rpcIP6
		}
	}

	fmt.Printf("rpc fingerprint: %x\n", fingerprint)
	fmt.Printf("rpc port:        %d\n", rpcPort)
	fmt.Printf("public key:      %x\n", publicKey)
	fmt.Printf("connect port:    %d\n", listenPort)
	fmt.Printf("IP4 IP6:         %s\n", IPs)

	fmt.Printf(txtRecord, IPs, listenPort, rpcPort, fingerprint, publicKey)
}

// extract first IP4 and/or IP6 connection
func getFirstConnections(connections []string) (string, string, int) {

	initialPort := 0
	IP4 := ""
	IP6 := ""

scan_connections:
	for i, c := range connections {
		if "" == c {
			continue scan_connections
		}
		v6, IP, port, err := splitConnection(c)
		if nil != err {
			exitwithstatus.Message("error: cannot decode[%d]: %q  error: %s", i, c, err)
		}
		if v6 {
			if "" == IP6 {
				IP6 = IP
				if 0 == initialPort || port == initialPort {
					initialPort = port
				}
			}
		} else {
			if "" == IP4 {
				IP4 = IP
				if 0 == initialPort || port == initialPort {
					initialPort = port
				}
			}
		}
	}
	return IP4, IP6, initialPort
}

// split connection into ip and port
func splitConnection(hostPort string) (bool, string, int, error) {
	host, port, err := net.SplitHostPort(hostPort)
	if nil != err {
		return false, "", 0, fault.ErrInvalidIpAddress
	}

	IP := net.ParseIP(strings.Trim(host, " "))
	if nil == IP {
		return false, "", 0, fault.ErrInvalidIpAddress
	}

	numericPort, err := strconv.Atoi(strings.Trim(port, " "))
	if nil != err {
		return false, "", 0, err
	}
	if numericPort < 1 || numericPort > 65535 {
		return false, "", 0, fault.ErrInvalidPortNumber
	}

	if nil != IP.To4() {
		return false, IP.String(), numericPort, nil
	}
	return true, "[" + IP.String() + "]", numericPort, nil
}

// get the working directory; if not set in the arguments
// it's set to the current directory
func getFilenameWithDirectory(arg map[string]string, name string) string {
	dir := "."
	if d, ok := arg["d"]; ok {
		dir = d
	}

	return filepath.Join(dir, name)
}

func makeSigningKey(version string, testnet bool, fileName string) error {
	var seed string
	var err error

	switch version {
	case "v1":
		seed, err = generateEncodedSeedV1(testnet)
	case "v2":
		seed, err = generateEncodedSeedV2(testnet)
	default:
		return fmt.Errorf("unsupported seed version")
	}

	if nil != err {
		return err
	}

	data := "SEED:" + seed + "\n"
	if err = ioutil.WriteFile(fileName, []byte(data), 0600); nil != err {
		return fmt.Errorf("error writing signing key file error: %s", err)
	}

	return nil
}

func generateEncodedSeedV1(testnet bool) (string, error) {
	sk := make([]byte, account.SecretKeyV1Length)
	if _, err := rand.Read(sk); nil != err {
		return "", fmt.Errorf("generate random secret key error: %s", err)
	}

	seed := make([]byte, 0)
	seed = append(seed, account.SeedHeaderV1...)
	if testnet {
		seed = append(seed, 0x01)
	} else {
		seed = append(seed, 0x00)
	}

	seed = append(seed, sk...)
	digest := sha3.Sum256(seed)
	checksum := digest[:account.SeedChecksumLength]
	seed = append(seed, checksum...)
	return util.ToBase58(seed), nil
}

func generateEncodedSeedV2(testnet bool) (string, error) {

	// space for 128 bit, extend to 132 bit later
	sk := make([]byte, 16, account.SecretKeyV2Length)

	n, err := rand.Read(sk)
	if nil != err {
		return "", err
	}

	if 16 != n {
		return "", fmt.Errorf("got %d bytes, expected is 16 bytes", n)
	}

	// extend to 132 bits
	sk = append(sk, sk[15]&0xf0)

	if account.SecretKeyV2Length != len(sk) {
		return "", fmt.Errorf("actual seed length is %d bytes, expected is %d bytes", len(sk), account.SecretKeyV2Length)
	}

	// network flag
	mode := sk[0]&0x80 | sk[1]&0x40 | sk[2]&0x20 | sk[3]&0x10
	if testnet {
		mode = mode ^ 0xf0
	}
	sk[15] = mode | sk[15]&0x0f

	// encode seed to base58
	seed := make([]byte, 0)
	seed = append(seed, account.SeedHeaderV2...)
	seed = append(seed, sk...)
	digest := sha3.Sum256(seed)
	checksum := digest[:account.SeedChecksumLength]
	seed = append(seed, checksum...)
	return util.ToBase58(seed), nil
}
