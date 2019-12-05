package main

import "math/big"

type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork{
	pow := ProofOfWork{
		block : block,
	}
	//指定难度值
	targeStr := "0001000000000000000000000000000000000000000000000000000000000000"
	//引入辅助变量
	tmpInt := big.Int{}
	//将难度值赋值给big.Int,使用16进制格式
	tmpInt.SetString(targeStr,16)
	pow.target = &tmpInt
	return &pow
}
