package main

import (
	"os"
	"fmt"
	"strconv"
)

//这是一个用来接收命令行参数并且控制区块链操作的文件
type CLI struct {
	bc *BlockChain
}

const Usage = `
	addBlock --data Data		"add data to blockChain"
	printChain					"正向打印区块链"
	printChainR					"反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"
	send FROM TO AMOUNT MINER DATA "由FROM转AMOUNT给TO，由MINER挖矿，同时写入DATA"
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
			//正向打印区块
			cli.PrintBlockChain()
		case "printChainR":
			//反向打印区块
			cli.PrintBlockChainReverse()
		case "getBalance":
			//获取余额
			if len(args) == 4 && args[2] == "--address" {
				address := args[3]
				cli.GetBalance(address)
			}
		case "send":
			if len(args) != 7 {
				fmt.Println("参数个数错误，请检查！")
				fmt.Println(Usage)
				return
			}
			fmt.Println("转账开始...")
			from := args[2]
			to := args[3]
			amount, _ := strconv.ParseFloat(args[4], 64)
			miner := args[5]
			data := args[6]
			cli.Send(from, to, amount, miner, data)
		default:
			fmt.Println(Usage)
	}
}

