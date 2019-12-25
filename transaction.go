package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"fmt"
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

func (tx *Transaction) IsCoinbase() bool {
	//1.交易input只有一个
	//if len(tx.TXInputs) == 1 {
	//	input := tx.TXInputs[0]
	//	//2.交易id为空
	//	//3/交易的index为-1
	//	if !bytes.Equal(input.TXid, []byte{}) || input.Index != -1 {
	//		return false
	//	}
	//}
	if len(tx.TXInputs) == 1 && len(tx.TXInputs[0].TXid) == 0 && tx.TXInputs[0].Index == -1 {
		return true
	}
	return false
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
	//对于挖矿交易，只有一个input和一个output
	tx := Transaction{[]byte{},[]TXInput{input},[]TXOutput{output}}
	tx.SetHash()

	return &tx
}

//创建普通的转账交易
func NewTransaction(from, to string, amount float64,bc *BlockChain) *Transaction {
	//1.找到最合理UTXO集合 map[string][]int64
	utxos, resValue := bc.FindNeedUTXOs(from, amount)

	if resValue < amount {
		fmt.Println("余额不足，交易失败！")
		return nil
	}

	var inputs []TXInput
	var outputs []TXOutput

	//2.将UTXO逐一转成input
	for id, indexArray := range utxos {
		for _, i := range indexArray {
			input := TXInput{[]byte(id), int64(i), from}
			inputs = append(inputs, input)
		}
	}
	//3.创建output
	output := TXOutput{amount, to}
	outputs = append(outputs, output)

	if resValue > amount {
		//4.如果有零钱，要找零
		outputs = append(outputs, TXOutput{resValue- amount,from})
	}
	tx := Transaction{[]byte{}, inputs, outputs}
	tx.SetHash()
	return &tx
}
