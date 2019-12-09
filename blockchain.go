package main

import (
	"gocode/20190724/go-bili/basecoin/bolt"
	"fmt"
)

//4.引入区块链
type BlockChain struct {
	//定义一个区块链数组
	//blocks []*Block
	db *bolt.DB

	// 存储最后一个区块的哈希
	tail []byte
}

const blockChainDb = "D:/lubo/gowork/src/gocode/20190724/go-bili/basecoin/blockChain.db"
const blockBucket = "blockBucket"
//5.定义一个区块链
func NewBlockChain() *BlockChain {

	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希，从数据库中读出lastHash
	var lastHash []byte

	db, err := bolt.Open(blockChainDb,0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			//创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash,genesisBlock.toByte())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			lastHash = genesisBlock.Hash
		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}
//定义一个创世块
func GenesisBlock() *Block{
	return NewBlock("Go一期创世块。", []byte{})
}
//6.添加区块
func (bc *BlockChain) AddBlock(data string){
	/*//获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks) - 1]
	prevHash := lastBlock.Hash

	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks,block)*/
}
