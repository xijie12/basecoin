package main

import (
	//"crypto/sha256"
	"time"
	"encoding/binary"
	"bytes"
	"log"
	"encoding/gob"
	"fmt"
)

//0.定义结构
type Block struct {
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.Merkel根（梅克尔根，这就是一个hash值）
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数，挖矿要找的数据
	Nonce uint64

	//a.当前区块哈希 正常比特币区块中没有当前区块的哈希，为了方便做了简化
	Hash []byte
	//b.数据
	Data []byte
}

func Uint64ToByte(num uint64) []byte{
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block{
	block := Block{
		Version: 1,
		PrevHash: prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp: uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce: 0,
		Hash: []byte{},
		Data: []byte(data),
	}

	//block.SetHash()
	//创建一个pow对象
	pow := NewProofOfWork(&block)
	//查找随机数，不停的进行哈希运算
	hash, nonce := pow.Run()
	//根据挖矿结果对区块数据进行更新
	block.Hash, block.Nonce = hash, nonce

	return &block
}

//序列化
func (block *Block) Serialize() []byte {
	var buffer bytes.Buffer
	//使用gob进行序列化（编码）得到字节流
	//1.定义一个编码器
	//2.使用编码器进行编码
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(&block)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("编码之后：%v\n", buffer.Bytes())
	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) Block {

	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&block)
	if err != nil {
		log.Fatal("decode:", err)
	}

	return block
}

/*//3.生成哈希
func (block *Block) SetHash() {
	//1.设置数据
	var blockInfo []byte
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.Hash...)
	blockInfo = append(blockInfo, block.Data...)

	tmp := [][]byte{
		Uint64ToByte(block.Version),
		block.PrevHash,
		block.MerkelRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Hash,
		block.Data,
	}
	//将二维切片数组连接起来，返回一个一维的切片
	blockInfo := bytes.Join(tmp, []byte{})

	//2.sha256
	//(一)
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
	//（二）
	myHash := sha256.New()
	bData := []byte(data)
	myHash.Write(bData)
	res := myHash.Sum(nil)
	block.Hash = res
}*/