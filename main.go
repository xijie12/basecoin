package main

import "fmt"

func main(){
	bc := NewBlockChain()
	bc.AddBlock("A向B转让50枚比特币！")

	it := bc.NewIterator()
	//调用迭代器，返回每一个区块的数据
	for {
		block := it.Next()

		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束！")
			break
		}
	}

	/*for i, block := range bc.blocks {
		fmt.Printf("=========当前区块高度: %d =========\n", i)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)
	}*/
}
