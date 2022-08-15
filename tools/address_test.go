package tools

import (
	"testing"
)

func TestBTCAddress1(t *testing.T) {
	add := "3G3BTrsFtvvUK8NJm3CMnsr2ARmGtHRQbq"
	t.Log(ValidateAddress(add))
}

func TestBTCAddress2(t *testing.T) {
	add := "1DUDSonGPkocRLUpjusvBtQBpJu69DzUQT"
	t.Log(ValidateAddress(add))
}

func TestBTCAddress3(t *testing.T) {
	add := "3GxbcKu1yuGMEnM6e32RZeFuMs1myVbxDU"
	t.Log(ValidateAddress(add))
}

func TestBTCAddress4(t *testing.T) {
	add := "bc1qex0aqq8mxqfh4cpl62eg755836djjx20yzuuu8"
	t.Log(ValidateAddress(add))
	t.Log(ValidateAddress("mgxif2Zb85XGwYzPkT2M76yKcxzubXdgWW"))
	t.Log(ValidateAddress("tb1qt7vjs5knzcn2r2tla03cafu4dm8ntkvr2q0dz0"))
	t.Log(ValidateAddress("mkHS9ne12qx9pS9VojpwU5xtRd4T7X7ZUt"))
	t.Log(ValidateAddress("2N3oefVeg6stiTb5Kh3ozCSkaqmx91FDbsm"))
	t.Log(ValidateAddress("92Pg46rUhgTT7romnV7iGW6W1gbGdeezqdbJCzShkCsYNzyyNcc"))
	t.Log(ValidateAddress("cNJFgo1driFnPcBdBX8BrJrpxchBWXwXCvNH5SoSkdcF6JXXwHMm"))
}

func TestEmailAddress1(t *testing.T) {
	t.Log(ValidateEmail(""))        // false
	t.Log(ValidateEmail("a@1.com")) // true
	t.Log(ValidateEmail("a@b.com")) // true
	t.Log(ValidateEmail("1@b.com")) // true
	t.Log(ValidateEmail("1@.com"))  // false
}

func TestEthAddress(t *testing.T) {
	t.Log(ValidateEthAddress(""))                                           // false
	t.Log(ValidateEthAddress("0xc1912fEE45d61C87Cc5EA59DaE31190FFFFf232d")) // true
	t.Log(ValidateEthAddress("0xc1912fee45d61c87cc5ea59dae31190fffff2323")) // true
	t.Log(ValidateEthAddress("0xc1912fee45d61c87cc5ea59dae31190fffff232G")) // false
	t.Log(ValidateEthAddress("0x0c1912fee45d61c87cc5ea59dae31190fffff232")) // true
	t.Log(ValidateEthAddress("0xC1912fEE45d61C87Cc5EA59DaE31190FFFFf232d")) // false
	t.Log(ValidateEthAddress("0xc1912fee45d61c87cc5ea59dae31190fffff232d")) // true
	t.Log(ValidateEthAddress("0xc1912fEE45d61C87Cc5EA59DaE31190FFFFf232d")) // true
}
