package main

import "testing"

func Test_DeriveAddress(t *testing.T) {
	// This is the hex gotten from my redeem script
	redeemScriptHex := "a82077c1468f4d4be46debf35472b3df7eb7c309ab24a78e6f209a4ee89336e330fb87"
	result := DeriveAddress(redeemScriptHex)

	if result != "2N2SehHypjCHZxaj5XxpLh3rdSUiTLMXt9H" {
		t.Errorf("Expected 2N2SehHypjCHZxaj5XxpLh3rdSUiTLMXt9H, got %v", result)
	}
}
