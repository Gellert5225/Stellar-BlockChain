package main

import (
	"fmt"
	"stellar_blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain()
	bc.AddBlock("Send 1 STLR to Gellert")
	bc.AddBlock("Send 5 STLR to Gellert")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
