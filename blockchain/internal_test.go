// Copyright (c) 2013-2017 The btcsuite developers
// Copyright (c) 2017 BitGo
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
This test file is part of the blockchain package rather than than the
blockchain_test package so it can bridge access to the internals to properly
test cases which are either not possible or can't reliably be tested via the
public interface.  The functions are only exported while the tests are being
run.
*/

package blockchain

import (
	"sort"
)

// TstSetCoinbaseMaturity makes the ability to set the coinbase maturity
// available to the test package.
func (b *BlockChain) TstSetCoinbaseMaturity(maturity uint16) {
	b.chainParams.CoinbaseMaturity = maturity
}

// TstTimeSorter makes the internal timeSorter type available to the test
// package.
func TstTimeSorter(times []int64) sort.Interface {
	return timeSorter(times)
}

// TstSetMaxMedianTimeEntries makes the ability to set the maximum number of
// median time entries available to the test package.
func TstSetMaxMedianTimeEntries(val int) {
	maxMedianTimeEntries = val
}

// TstCheckBlockScripts makes the internal checkBlockScripts function available
// to the test package.
var TstCheckBlockScripts = checkBlockScripts

// TstDeserializeUtxoEntry makes the internal deserializeUtxoEntry function
// available to the test package.
var TstDeserializeUtxoEntry = deserializeUtxoEntry
