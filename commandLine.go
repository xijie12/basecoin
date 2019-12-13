package main

import "fmt"

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("添加区块成功！")
}

func (cli *CLI) PrintBlockChain() {
	bc := cli.bc
	it := bc.NewIterator()
	//调用迭代器，返回每一个区块的数据
	for {
		block := it.Next()
		fmt.Println("=======================================")
		fmt.Printf("版本号： %d\n", block.Version)
		fmt.Printf("前区块哈希值： %x\n", block.PrevHash)
		fmt.Printf("梅克尔根： %x\n", block.MerkelRoot)
		fmt.Printf("时间戳： %d\n", block.TimeStamp)
		fmt.Printf("难度值： %d\n", block.Difficulty)
		fmt.Printf("随机数： %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("区块数据： %s\n", block.Data)

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链遍历结束！")
			break
		}
	}
}
