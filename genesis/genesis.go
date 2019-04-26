// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package genesis

import (
	"github.com/bitmark-inc/bitmarkd/blockdigest"
)

// the starting block number
const (
	BlockNumber = uint64(1)
)

// this is block 1, the Genesis Block
// ----------------------------------

// LIVE Network
// ------------

// LiveGenesisBlock - live genesis block data
var LiveGenesisBlock = []byte{
	0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x63, 0x8c, 0x15, 0x9c,
	0x1f, 0x11, 0x3f, 0x70, 0xa9, 0x86, 0x6d, 0x9a,
	0x9e, 0x52, 0xe9, 0xef, 0xe9, 0xb9, 0x92, 0x08,
	0x48, 0xad, 0x1d, 0xf3, 0x48, 0x51, 0xbe, 0x8a,
	0x56, 0x2a, 0x99, 0x8d, 0xb7, 0x9a, 0x80, 0x56,
	0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x00, 0x11, 0x5a, 0x38, 0xbf,
	0x3a, 0x90, 0x9f, 0xe1, 0x01, 0x00, 0x14, 0x44,
	0x4f, 0x57, 0x4e, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x52, 0x41, 0x42, 0x42, 0x49, 0x54, 0x20, 0x68,
	0x6f, 0x6c, 0x65, 0x21, 0x11, 0x4a, 0x65, 0xf1,
	0xd2, 0x06, 0x50, 0x08, 0x12, 0x76, 0xf0, 0x1d,
	0xf4, 0x3e, 0x70, 0x55, 0x4e, 0x95, 0x49, 0x8f,
	0x37, 0x78, 0xe5, 0x6d, 0xaa, 0x2c, 0x49, 0x82,
	0x03, 0xae, 0x9c, 0x70, 0xe6, 0xf4, 0xca, 0xb9,
	0xd2, 0xd2, 0xcc, 0xdd, 0xb4, 0x4c, 0x40, 0xc2,
	0xa3, 0x84, 0xeb, 0xc9, 0x01, 0xa1, 0x8a, 0x13,
	0xa2, 0x70, 0xaa, 0x9f, 0x5e, 0x08, 0x06, 0x77,
	0xd7, 0xab, 0x2f, 0xd8, 0x88, 0xa5, 0xf6, 0x57,
	0xd2, 0xc6, 0xd4, 0x69, 0x2e, 0x6f, 0xcd, 0xe7,
	0x1c, 0x04, 0xb9, 0x1b, 0xe1, 0x40, 0x0e, 0x7c,
	0x1e, 0x8d, 0x5e, 0x2b, 0x34, 0x83, 0xc4, 0x77,
	0xfe, 0xa1, 0x7b, 0xc1, 0xde, 0xe0, 0x05, 0xcc,
	0x8d, 0x4d, 0xf8, 0x62, 0x77, 0x0d, 0x0c,
}

// LiveGenesisDigest - digest of the live genesis header
// 0012af8f437cc0b4358df93f317129747ab3cf3cf0795530decd01eb39f7935c
var LiveGenesisDigest = blockdigest.Digest([...]byte{
	0x5c, 0x93, 0xf7, 0x39, 0xeb, 0x01, 0xcd, 0xde,
	0x30, 0x55, 0x79, 0xf0, 0x3c, 0xcf, 0xb3, 0x7a,
	0x74, 0x29, 0x71, 0x31, 0x3f, 0xf9, 0x8d, 0x35,
	0xb4, 0xc0, 0x7c, 0x43, 0x8f, 0xaf, 0x12, 0x00,
})

// TEST Network
// ------------

// TestGenesisBlock - testnet genesis block data
var TestGenesisBlock = []byte{
	0x01, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0xee, 0x07, 0xbb, 0xc3,
	0xd7, 0x49, 0xe0, 0x7d, 0x24, 0xb9, 0x0c, 0xd1,
	0xec, 0x35, 0x14, 0x70, 0x2e, 0x87, 0x85, 0x22,
	0xda, 0xf7, 0x16, 0xc1, 0x73, 0x24, 0xd6, 0x66,
	0x69, 0x7b, 0x8a, 0x63, 0x4b, 0x42, 0x78, 0x54,
	0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x00, 0xd4, 0x4c, 0x2b, 0xca,
	0xee, 0x40, 0x36, 0x47, 0x01, 0x00, 0x1d, 0x42,
	0x69, 0x74, 0x6d, 0x61, 0x72, 0x6b, 0x20, 0x54,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x47,
	0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x20, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x21, 0x13, 0xb2, 0xb5,
	0x04, 0x82, 0x7f, 0x30, 0xa8, 0xdc, 0x1b, 0x75,
	0x95, 0xeb, 0xb9, 0x88, 0xdc, 0xf8, 0x7c, 0xad,
	0xac, 0x9e, 0x3a, 0x38, 0xf6, 0xbe, 0x81, 0x8c,
	0x72, 0xbe, 0x03, 0x35, 0xfa, 0x74, 0xf4, 0xca,
	0xb9, 0xd2, 0xc2, 0xee, 0xdc, 0xb2, 0x54, 0x40,
	0x02, 0xa8, 0xbf, 0x5c, 0x21, 0x73, 0x03, 0x24,
	0x04, 0x40, 0x79, 0xa5, 0x78, 0x0a, 0x9c, 0xd2,
	0x2f, 0xc2, 0x22, 0xb4, 0x4c, 0x91, 0x29, 0x17,
	0xce, 0xa5, 0xb9, 0xd3, 0x77, 0x0c, 0x13, 0x8e,
	0x8d, 0x3e, 0xae, 0x98, 0xb7, 0x6c, 0x2e, 0x93,
	0xa9, 0x7e, 0x41, 0xc4, 0x1b, 0xae, 0x36, 0xc8,
	0x41, 0x37, 0x08, 0xa9, 0x94, 0xfe, 0xc2, 0xf9,
	0xeb, 0xc0, 0xf8, 0x02, 0x98, 0x3d, 0xf6, 0x01,
}

// TestGenesisDigest - digest of the test genesis header
// 00fe807b8f2c5a5416bc570b289baadebb9596daeec5128f6b924a7cb88be68a
var TestGenesisDigest = blockdigest.Digest([...]byte{
	0x8a, 0xe6, 0x8b, 0xb8, 0x7c, 0x4a, 0x92, 0x6b,
	0x8f, 0x12, 0xc5, 0xee, 0xda, 0x96, 0x95, 0xbb,
	0xde, 0xaa, 0x9b, 0x28, 0x0b, 0x57, 0xbc, 0x16,
	0x54, 0x5a, 0x2c, 0x8f, 0x7b, 0x80, 0xfe, 0x00,
})
