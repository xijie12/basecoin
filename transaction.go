package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

const reward = 12.5

//1.定义交易结构
type Transaction struct {
	TXID      []byte		//交易ID
	TXInputs  []TXInput		//交易输入数组
	TXOutputs []TXOutput	//交易输出的数组
}
//交易输入
type TXInput struct {
	TXid []byte	//引用交易ID
	Index int64	//引用的output的索引值
	Sig string	//解锁脚本，用地址来模拟
}
//交易输出
type TXOutput struct {
	Value float64	//转账金额
	PubKeyHash string	//锁定脚本，用地址模拟
}
//设置交易ID
func (tx *Transaction) SetHash() {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	data := buffer.Bytes()
	hash := sha256.Sum256(data)
	tx.TXID = hash[:]
}
//2.提供创建交易方法(挖矿交易）
func NewCoinbaseTX(address string, data string) *Transaction {
	//挖矿交易的特点：
	//1.只有一个input
	//2.无需引用交易id
	//3.无需引用index
	//矿工由于挖矿时无需指定签名，所以sig字段可以由矿工自由填写数据，一般填写矿池名字
	input := TXInput{[]byte{},-1,data}
	output := TXOutput{reward, address}
	tx := Transaction{[]byte{},[]TXInput{input},[]TXOutput{output}}
	tx.SetHash()

	return &tx
}
//3.创建挖矿交易
//4.根据交易调整程序
