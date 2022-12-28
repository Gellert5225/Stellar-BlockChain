package main

import (
	blck "stellar_blockchain/blockchain"
	"stellar_blockchain/cli"
)

func main() {
	bc := blck.NewBlockChain()
	defer bc.DB.Close()

	cli := cli.CLI{BC: bc}
	cli.Run()
}
