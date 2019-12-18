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
	printChain					"正向打印区块链"
	printChainR					"反向打印区块链"
	getBalance --address ADDRESS "获取指定地址的余额"
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

		default:
			fmt.Println(Usage)
	}
}

func (cli * CLI) GetBalance(address string) {
	utxos := cli.bc.FindUTXOs(address)
	total := 0.0
	for _,utxo := range utxos {
		total += utxo.Value
	}

	fmt.Printf("\"%s\" 的余额为: %f\n", address, total)
}