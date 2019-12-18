package main

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
//2.提供创建交易方法
//3.创建挖矿交易
//4.根据交易调整程序
