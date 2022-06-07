package main

import (
	"blockchain/public_blockchain/part4-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	// 创世区块
	blockchain := BLC.CreateGenesisBlockchainGenesisBlock()
	fmt.Println(blockchain)
	//新区块
	blockchain.AddBlockToBlockchain("send 100 to Tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("send 200 to Tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("send 300 to Tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("send 400 to Tom", blockchain.Blocks[len(blockchain.Blocks)-1].Height+1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	fmt.Println(blockchain)

}
