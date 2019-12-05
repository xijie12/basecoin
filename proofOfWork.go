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
	targeStr := "0001000000000000000000000000000000000000000000000000000000000000"
	tmpInt := big.Int{}
	tmpInt.SetString(targeStr,16)
	pow.target = &tmpInt
	return &pow
}
