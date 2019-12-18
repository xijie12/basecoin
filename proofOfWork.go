package main

import (
	"math/big"
	"bytes"
	"crypto/sha256"
	"fmt"
)

//1.声明一个结构体
type ProofOfWork struct {
	block *Block
	target *big.Int
}

//2.创建一个工作量证明
func NewProofOfWork(block *Block) *ProofOfWork{
	pow := ProofOfWork{
		block : block,
	}
	//指定难度值
	targeStr := "0000100000000000000000000000000000000000000000000000000000000000"
	//引入辅助变量
	tmpInt := big.Int{}
	//将难度值赋值给big.Int,使用16进制格式
	tmpInt.SetString(targeStr,16)
	pow.target = &tmpInt
	return &pow
}

//3.提供不断计算hash
func (pow *ProofOfWork) Run() ([]byte, uint64){
	var nonce uint64
	block := pow.block
	var hash [32]byte

	fmt.Println("开始挖矿...")
	for {
		//1.拼装数据（区块的数据，还有不断变化的随机数）
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Hash,
			//只对区块头做哈希值，区块体通过MerkelRoot产生影响
			//block.Data,
		}
		//将二维切片数组连接起来，返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})
		//2.做hash运算
		//func Sum256(data []byte) [size]byte {}
		hash = sha256.Sum256(blockInfo)
		//3.与pow中的target进行比较
		tmpInt := big.Int{}
		//将得到的hash数组转换成一个big.Int
		tmpInt.SetBytes(hash[:])
		//比较当前的hash与目标hash，如果当前的哈希值小于目标的哈希值，就说明找到了，否则继续查找
		// -1 if x < y
		// 0 if x == y
		// +1 if x > y
		//a.找到了，返回退出
		//b.没找到，继续找，随机数加1
		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功！hash: %x, nonce: %d\n", hash, nonce)
			break
		} else {
			nonce++
		}

	}
	return hash[:], nonce
}

//4.提供一个校验函数
func IsValid(){

}