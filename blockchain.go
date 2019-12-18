package main

import (
	"gocode/20190724/go-bili/basecoin/bolt"
	"fmt"
	"log"
	"bytes"
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
func NewBlockChain(address string) *BlockChain {

	//return &BlockChain{
	//	blocks: []*Block{genesisBlock},
	//}

	//最后一个区块的哈希，从数据库中读出lastHash
	var lastHash []byte

	db, err := bolt.Open(blockChainDb,0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	//defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			//创建一个创世块，并作为第一个区块添加到区块链中
			genesisBlock := GenesisBlock(address)
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte("lastHashKey"),genesisBlock.Hash)
			lastHash = genesisBlock.Hash

			//测试0
			//blockBytes := bucket.Get(genesisBlock.Hash)
			//block := Deserialize(blockBytes)
			//fmt.Printf("block info: %v\n", block)

		} else {
			lastHash = bucket.Get([]byte("lastHashKey"))
		}

		return nil
	})

	return &BlockChain{db, lastHash}
}
//定义一个创世块
func GenesisBlock(address string) *Block{
	coinbase := NewCoinbaseTX(address, "Go一期创世块。")
	return NewBlock([]*Transaction{coinbase}, []byte{})
}
//6.添加区块
func (bc *BlockChain) AddBlock(txs []*Transaction){
	//获取前一个区块的哈希
	db := bc.db
	lastHash := bc.tail
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("bucket 不存在，请检查！")
		}

		//创建一个区块
		block := NewBlock(txs, lastHash)

		//更新bolt上一个区块哈希
		bucket.Put(block.Hash,block.Serialize())
		bucket.Put([]byte("lastHashKey"),block.Hash)

		//更新内存中的上一个区块哈希
		bc.tail = block.Hash

		return nil
	})
}

func (bc *BlockChain) PrintChain() {
	blockHeight := 0
	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("blockBucket"))
		b.ForEach(func(k, v []byte) error {
			if bytes.Equal(k, []byte("LastHashKey")){
				return nil
			}
			block := Deserialize(v)
			fmt.Printf("===================区块链高度：%d====================\n", blockHeight)
			blockHeight++
			fmt.Printf("版本号： %d\n", block.Version)
			fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
			fmt.Printf("梅克尔根： %x\n", block.MerkelRoot)
			fmt.Printf("时间戳： %d\n", block.TimeStamp)
			fmt.Printf("难度值： %d\n", block.Difficulty)
			fmt.Printf("随机数： %d\n", block.Nonce)
			fmt.Printf("当前区块哈希值： %x\n", block.Hash)
			fmt.Printf("区块数据： %s\n", block.Transactions[0].TXInputs[0].Sig)
			return nil
		})
		return nil
	})
}

//找到指定地址的所有UTXO
func (bc *BlockChain) FindUTXOs(address string) []TXOutput {
	var UTXO []TXOutput

	return UTXO
}
