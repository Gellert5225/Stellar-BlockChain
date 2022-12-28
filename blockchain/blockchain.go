package blck

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type BlockChain struct {
	head []byte
	DB   *bolt.DB
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (blockChain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	// read only transaction
	err := blockChain.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = blockChain.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		blockChain.head = newBlock.Hash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}

func NewBlockChain() *BlockChain {
	var head []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	// open a read-write transaction Tx
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("No existing blockchain found, creating new one.")
			genesis := NewGenesisBlock()

			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// "l" is the hash of the last block
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}

			head = genesis.Hash
		} else {
			head = b.Get([]byte("l"))
		}

		return nil
	})

	bc := BlockChain{head, db}

	return &bc
}

func (blockChain *BlockChain) Iterator() *BlockChainIterator {
	bci := &BlockChainIterator{blockChain.head, blockChain.DB}

	return bci
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = Deserialize(encodedBlock)

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
