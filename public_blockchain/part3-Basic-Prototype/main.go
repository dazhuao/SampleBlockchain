package main

import (
	"blockchain/public_blockchain/part3-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	genesisBlockchain := BLC.CreateGenesisBlockchainGenesisBlock()
	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])
}
