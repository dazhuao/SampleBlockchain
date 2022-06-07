package BLC

type Blockchain struct {
	Blocks []*Block
}

// 创建带有创世区块的区块链
func CreateGenesisBlockchainGenesisBlock() *Blockchain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data")
	//返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}

// 增加区块到区块链中
func (blc *Blockchain) AddBlockToBlockchain(data string,height int64,preHash []byte) {
	// 创建新区块
	newBlock := NewBlock(data,height,preHash)
	// 在链上添加区块
	blc.Blocks = append(blc.Blocks, newBlock)
}