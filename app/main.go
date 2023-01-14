package main

import (
	"fmt"
	"strconv"

	"blockchaintest/internal/blockchain"
)

func main() {
	blockChain := blockchain.InitBlockChain()

	blockChain.AddBlock("first block")
	blockChain.AddBlock("second block")
	blockChain.AddBlock("third block")

	for _, block := range blockChain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)

		proof := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(proof.Validate()))
		fmt.Println()
	}
}
