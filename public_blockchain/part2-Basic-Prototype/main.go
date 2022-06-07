package main

import (
	"blockchain/public_blockchain/part2-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	genesisBlock := BLC.CreateGenesisBlock("genesis block")
	fmt.Println(genesisBlock)
}
