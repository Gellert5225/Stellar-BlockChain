package blck

type BlockChain struct {
	Blocks []*Block
}

func (blockChain *BlockChain) AddBlock(data string) {
	prev_block := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock := NewBlock(data, prev_block.Hash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
