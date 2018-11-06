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
	"testing"
)

func Test1(t *testing.T) {
	BlockID("0501020304050607080102030405060708010203040506070801020304050607080102030405060708010203040506070801020304050607080102030405060708")
}
func Test2(t *testing.T) {
	BlockID("00")
}
func Test3(t *testing.T) {
	BlockID("0500000000000000000000000000000000000000000000000000000000000000080102030405060708010203040506070801020304050607080102030405060708")
}
