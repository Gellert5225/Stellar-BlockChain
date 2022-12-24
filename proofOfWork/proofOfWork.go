package proofOfWork

import (
	"bytes"
	"math/big"
	"stellar_blockchain/block"
)

const TARGET_BITS = 24

type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))

	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(TARGET_BITS)),
		IntToHex(int64(nonce)),
	}, []byte{})
	return data
}
