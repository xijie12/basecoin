package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"github.com/btcsuite/btcutil/base58"
)

//这里的钱包是—结构，每一个钱包保存了公钥，私钥对

type Wallet struct {
	//私钥
	Private *ecdsa.PrivateKey
	//PubKey *ecdsa.PublicKey
	//这里的PubKey不存储原始的公钥，而是存储x和y拼接的字符串，在校验端重新拆分（参考r, s）
	PubKey []byte
}

//创建钱包
func NewWallet() *Wallet{
	//创建曲线
	curve := elliptic.P256()
	//生成私钥
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	//生成公钥
	pubKeyOrig := privateKey.PublicKey
	pubKey := append(pubKeyOrig.X.Bytes(), pubKeyOrig.Y.Bytes()...)

	return &Wallet{privateKey, pubKey}
}

//生成地址
func (w *Wallet)NewAddress() string{

	pubKey := w.PubKey

	rip160HashValue := HashPubKey(pubKey)

	version := byte(00)
	//拼接version
	payload := append([]byte{version}, rip160HashValue...)

	checkCode := CheckSum(payload)

	//25字节数据
	payload = append(payload, checkCode...)

	//btcd,这个是比特币全节点源码
	address := base58.Encode(payload)
	return address
}

func HashPubKey(data []byte) []byte {

	hash := sha256.Sum256(data)

	rip160hasher := ripemd160.New()
	_, err := rip160hasher.Write(hash[:])
	if err != nil {
		log.Panic(err)
	}
	//返回rip160的哈希结果
	rip160HashValue := rip160hasher.Sum(nil)

	return rip160HashValue
}

func CheckSum(data []byte) []byte{
	//checksum
	hash1 := sha256.Sum256(data)
	hash2 := sha256.Sum256(hash1[:])

	//前4字节校验
	checkCode := hash2[:4]

	return checkCode
}