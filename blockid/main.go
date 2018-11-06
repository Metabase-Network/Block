// Copyright 2014 The Metabase Authors
// This file is part of vasuki p2p library.
//
// The vasuki library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The vasuki library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package blockid

import (
	"encoding/hex"
	"math/big"
)

func BlockID(blockID string) (b Blockid) {
	bid, err := hex.DecodeString(blockID)
	if isValidsize(bid) && (err == nil) {
		if isValidchainID(bid[0:1], bid[1:33]) {
			b := Blockid{ENo: bid[0:1], CNo: bid[1:33], BNo: bid[34:66], state: true, raw: blockID}
			return b
		} else {
			b := Blockid{ENo: []byte("00"), CNo: []byte("00"), BNo: []byte("00"), state: false, raw: blockID}
			return b
		}
	} else {
		b := Blockid{ENo: []byte("00"), CNo: []byte("00"), BNo: []byte("00"), state: false, raw: blockID}
		return b
	}
}

func isValidsize(b []byte) bool {
	if len(b) == 65 {
		return true
	} else {
		return false
	}
}

func isValidchainID(Eno []byte, Cno []byte) bool {
	set2, _ := hex.DecodeString("02")
	set0, _ := hex.DecodeString("00")
	Con0 := new(big.Int).SetBytes(set0)
	Con2 := new(big.Int).SetBytes(set2)
	Epoch := new(big.Int).SetBytes(Eno)
	Chain := new(big.Int).SetBytes(Cno)

	ChainEpoch := Epoch.Exp(Con2, Epoch, Con0)
	ret := ChainEpoch.Cmp(Chain)
	if ret >= 0 {
		return true
	} else {
		return false
	}
}

func (b Blockid) FetchChainNo() *big.Int {
	return new(big.Int).SetBytes(b.CNo)
}

func (b Blockid) FetchEpochNo() *big.Int {
	return new(big.Int).SetBytes(b.ENo)
}
