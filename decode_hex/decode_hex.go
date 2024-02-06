package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"log"
	"strings"
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

func main() {
	rawTransactionHex := "020000000001010ccc140e766b5dbc884ea2d780c5e91e4eb77597ae64288a42575228b79e234900000000000000000002bd37060000000000225120245091249f4f29d30820e5f36e1e5d477dc3386144220bd6f35839e94de4b9cae81c00000000000016001416d31d7632aa17b3b316b813c0a3177f5b6150200140838a1f0f1ee607b54abf0a3f55792f6f8d09c3eb7a9fa46cd4976f2137ca2e3f4a901e314e1b827c3332d7e1865ffe1d7ff5f5d7576a9000f354487a09de44cd00000000"

	// Remove spaces and newline characters from the input
	rawTransactionHex = strings.ReplaceAll(rawTransactionHex, " ", "")
	rawTransactionHex = strings.ReplaceAll(rawTransactionHex, "\n", "")

	DecodeTransaction(rawTransactionHex)
}
