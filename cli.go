package main

import (
	"os"
	"fmt"
)

//这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data Data		"add data to blockChain"
	printChain					"print all blockChain data"
`

func (cli *CLI) Run(){
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}

	cmd := args[1]
	switch cmd {
		case "addBlock":
			//添加区块
			if len(args) == 4 && args[2] == "--data" {
				data := args[3]
				cli.AddBlock(data)
			}else{
				//fmt.Printf("添加区块参数使用不当，请检查")
				fmt.Printf(Usage)
			}
		case "printChain":
			//打印区块
			cli.PrintBlockChain()
		default:
			fmt.Println(Usage)
	}
}