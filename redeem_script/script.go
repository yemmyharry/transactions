package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetRedeemScript(preimage string) string {
	// Encode the preimage
	preimageBytes := []byte(preimage)

	// Compute the SHA-256 hash
	hasher := sha256.New()
	hasher.Write(preimageBytes)
	hashBytes := hasher.Sum(nil)

	// Encode the hash to hex
	hashHex := hex.EncodeToString(hashBytes)
	//OP_SHA256 = 0xa8
	//OP_EQUAL= 87
	redeemScript := "a8" + hashHex + "87"

	fmt.Println("Redeem script (hex):", redeemScript)

	return redeemScript
}
