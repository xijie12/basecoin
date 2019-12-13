package main

import (
	"gocode/20190724/go-bili/basecoin/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte //游标，用于不断索引
}

func (bc *BlockChain) NewIterator() *BlockChainIterator{

	return &BlockChainIterator{
		bc.db,
		bc.tail,	//最初指向区块链的最后一个区块，随着Next的调用，不断变化
	}
}

func (it *BlockChainIterator)Next() *Block{
	var block Block
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket == nil {
			log.Panic("迭代器遍历时bucket不应该为空，请检查！")
		}
		blockTmp := bucket.Get(it.currentHashPointer)
		block = Deserialize(blockTmp)
		//哈希左移
		it.currentHashPointer = block.PrevHash
		return nil
	})
	return &block
}