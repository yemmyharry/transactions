package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"log"
)

func DecodeTransaction(rawTransactionHex string) (txn wire.MsgTx) {
	// Decode raw transaction hex into a byte slice
	transactionBytes, err := hex.DecodeString(rawTransactionHex)
	if err != nil {
		log.Fatal("Error decoding raw transaction hex:", err)
	}

	// Deserialize the transaction
	var transaction wire.MsgTx
	err = transaction.Deserialize(bytes.NewReader(transactionBytes))
	if err != nil {
		log.Fatal("Error deserializing transaction:", err)
	}

	// Print version
	fmt.Println("Version:", transaction.Version)

	fmt.Println("\nInputs:")
	for _, input := range transaction.TxIn {
		fmt.Printf("  Transaction Hash : %s\n", input.PreviousOutPoint.Hash)
		fmt.Printf("  Output Index: %d\n", input.PreviousOutPoint.Index)
		fmt.Printf("  Script Signature: %v\n", input.SignatureScript)
		fmt.Printf("  Sequence %d \n", input.Sequence)

	}

	fmt.Println("\nOutputs:")
	for _, output := range transaction.TxOut {
		fmt.Printf("  Value: %d\n", output.Value)
		fmt.Printf("  Script Output: %x\n", output.PkScript)
		fmt.Printf("  Serialised size: %x\n", output.SerializeSize())
		fmt.Println("               ")

	}

	fmt.Println("\nLocktime:", transaction.LockTime)

	return transaction
}
