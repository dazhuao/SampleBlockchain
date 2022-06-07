package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Height        int64  // 1.区块高度（第几个区块）
	PrevBlockHash []byte // 2.上一个区块的hash值
	Data          []byte // 3.交易数据
	Timestamp     int64  // 4.时间戳
	Hash          []byte // 5.Hash值
}

func (block *Block) SetHash() {
	//1.Height 转化为字节数组
	heightBytes := IntToHex(block.Height)
	//2.Timestamp转化为字节数组
	timeString := strconv.FormatInt(block.Timestamp, 2)
	timeBytes := []byte(timeString)
	//3.拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
	//4.生成Hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

// 1.创建新的区块
func NewBlock(data string, height int64, prevchanBlockHash []byte) *Block {
	//创建区块
	block := &Block{
		Height:        height,
		PrevBlockHash: prevchanBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}
	//计算区块Hash
	block.SetHash()
	return block

}
// 生成创世区块
func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
