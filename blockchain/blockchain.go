package blockchain

import (
	"stellar_blockchain/block"
)

type BlockChain struct {
	Blocks []*block.Block
}

func (blockChain *BlockChain) AddBlock(data string) {
	prev_block := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock := block.NewBlock(data, prev_block.Hash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*block.Block{block.NewGenesisBlock()}}
}
