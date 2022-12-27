package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("Send 1 STLR to Gellert")
	bc.AddBlock("Send 5 STLR to Gellert")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
