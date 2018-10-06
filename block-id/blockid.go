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
	"errors"
	"math/big"
)

type id struct {
	epochID big.Int
	chainID big.Int
	blockNo big.Int
}

var errInvalidBlockID = errors.New("invalid BlockID")

func genBlockID(bid string) id {
	  var e,_ =new(big.Int).SetString(bid[0:4], 10)
  	  var c, _=new(big.Int).SetString(bid[4:68], 10)
 	  var b,_=new(big.Int).SetString(bid[68:132], 10)
 	  return id{epochID:*e, chainID:*c, blockNo:*b}
}
func Validate(bid string) bool{

  	 var ID id
 	 ID=GenBlockID(bid)//store blockid in struct
 	 //validation
 	 i:=big.NewInt(2)
 	 i.Exp(i,&ID.epochID,nil)
 	 if ID.chainID.Cmp(i)<1{
   	 return true
 }
  return false
}
}
