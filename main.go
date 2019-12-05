package main

import "fmt"

func main(){
	bc := NewBlockChain()
	bc.AddBlock("A向B转出50枚比特币！")
	for i, block := range bc.blocks {
		fmt.Printf("=========当前区块高度: %d =========\n", i)
		fmt.Printf("版本号： %d\n", block.Version)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("梅克尔根： %x\n", block.MerkelRoot)
		fmt.Printf("时间戳： %d\n", block.TimeStamp)
		fmt.Printf("时间戳： %d\n", block.Difficulty)
		fmt.Printf("随机数： %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}
}
