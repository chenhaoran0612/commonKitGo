package tools

import (
   "testing"
)

func TestUint256(t *testing.T){
   t.Log(AddUInt(1000,10000).Hex())
   t.Log(AddUInt(1000,10000).String())
   t.Log(AddUInt(1000,10000).Uint64())
   t.Log(MulUInt(1000,1000000000000000000).Uint64())
   uint64V, isOverflow := ToUInt256ToUInt(MulUInt(1000,1000000000000000000))
   t.Log(uint64V, isOverflow)
}
