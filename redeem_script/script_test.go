package main

import (
	"testing"
)

func Test_GetRedeemScript(t *testing.T) {

	preimage := "Btrust Builders"

	result := GetRedeemScript(preimage)

	if result != "a816e05614526c1ebd3a170a430a1906a6484fdd203ab7ce6690a54938f5c44d7d87" {
		t.Errorf("Expected a816e05614526c1ebd3a170a430a1906a6484fdd203ab7ce6690a54938f5c44d7d87, got %v", result)
	}

}
