package tools


import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/bech32"
)

func TestBech32Address(t *testing.T){
	encoded := "bc1qex0aqq8mxqfh4cpl62eg755836djjx20yzuuu8"
	hrp, decoded, err := bech32.Decode(encoded)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Show the decoded data.
	fmt.Println("Decoded human-readable part:", hrp)
	fmt.Println("Decoded Data:", hex.EncodeToString(decoded))
}