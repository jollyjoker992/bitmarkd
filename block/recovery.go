// Copyright (c) 2014-2019 Bitmark Inc.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package block

import (
	"encoding/binary"

	"github.com/bitmark-inc/bitmarkd/blockrecord"
	"github.com/bitmark-inc/bitmarkd/storage"
)

func doBlockHeaderHash() error {
	return storage.Pool.Blocks.NewFetchCursor().Map(recoverBlockHeaderHash)
}

func recoverBlockHeaderHash(blockNumberBytes []byte, packedBlock []byte) error {
	globalData.Lock()
	defer globalData.Unlock()

	// TODO: decide if we need to disable reservoir when migrating the block db
	// reservoir.Disable()
	// defer reservoir.Enable()

	// reservoir.ClearSpend()

	blockNumber := binary.BigEndian.Uint64(blockNumberBytes)

	blockHeaderHashBytes := storage.Pool.BlockHeaderHash.Get(blockNumberBytes)
	if blockHeaderHashBytes == nil {
		digest, err := blockrecord.ComputeHeaderHash(packedBlock)
		if nil != err {
			return err
		}

		storage.Pool.BlockHeaderHash.Put(blockNumberBytes, digest[:])
	}

	globalData.log.Debugf("rebuilt block: %d", blockNumber)

	return nil
}

func doRecovery() error {
	return storage.Pool.Blocks.NewFetchCursor().Map(recoverRecord)
}

func recoverRecord(key []byte, value []byte) error {
	return StoreIncoming(value, NoRescanVerified)
}
