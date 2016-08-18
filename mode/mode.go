// Copyright (c) 2014-2016 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package mode

import (
	"github.com/bitmark-inc/bitmarkd/chain"
	"github.com/bitmark-inc/bitmarkd/fault"
	"github.com/bitmark-inc/logger"
	"sync"
)

// type to hold the mode
type Mode int

// all possible modes
const (
	Stopped       = Mode(iota)
	Resynchronise = Mode(iota)
	Normal        = Mode(iota)
	maximum       = Mode(iota)
)

var globals struct {
	sync.RWMutex
	log     *logger.L
	mode    Mode
	testing bool
	chain   string
}

// set up the mode system
func Initialise(chainName string) error {

	// ensure start up in resynchronise mode
	globals.Lock()
	defer globals.Unlock()

	globals.log = logger.New("mode")
	globals.log.Info("starting…")

	// default settings
	globals.chain = chainName
	globals.testing = false
	globals.mode = Resynchronise

	// override for specific chain
	switch chainName {
	case chain.Bitmark:
		// no change
	case chain.Testing, chain.Local:
		globals.testing = true
	default:
		globals.log.Criticalf("mode cannot handle chain: '%s'", chainName)
		return fault.ErrInvalidChain
	}
	return nil
}

// shutdown mode handling
func Finalise() {
	Set(Stopped)
	globals.log.Info("shutting down…")
}

// change mode
func Set(mode Mode) {

	if mode >= Stopped && mode < maximum {
		globals.Lock()
		globals.mode = mode
		globals.Unlock()

		globals.log.Infof("set: %d", mode)
	} else {
		globals.log.Errorf("ignore invalid set: %d", mode)
	}
}

// detect mode
func Is(mode Mode) bool {
	globals.RLock()
	defer globals.RUnlock()
	return mode == globals.mode
}

// detect mode
func IsNot(mode Mode) bool {
	globals.RLock()
	defer globals.RUnlock()
	return mode != globals.mode
}

// special for testing
func IsTesting() bool {
	globals.RLock()
	defer globals.RUnlock()
	return globals.testing
}

// name of the current chain
func ChainName() string {
	globals.RLock()
	defer globals.RUnlock()
	return globals.chain
}

// current mode rep[resented as a string
func String() string {
	globals.RLock()
	defer globals.RUnlock()
	switch globals.mode {
	case Stopped:
		return "Stopped"
	case Resynchronise:
		return "Resynchronise"
	case Normal:
		return "Normal"
	default:
		return "*Unknown*"
	}
}
