package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"

	"log"
)

func DeriveAddress(redeemScriptHex string) string {
	redeemScript, err := hex.DecodeString(redeemScriptHex)
	if err != nil {
		log.Fatal(err)
	}

	// Derive the P2SH address from the redeem script on TestNet
	scriptHash := btcutil.Hash160(redeemScript)
	p2shAddress, err := btcutil.NewAddressScriptHash(scriptHash, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("P2SH Address:", p2shAddress.EncodeAddress())
	return p2shAddress.EncodeAddress()
}
