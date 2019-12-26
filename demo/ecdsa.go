package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"crypto/sha256"
	"log"
	"math/big"
)

//1.演示如何使用ecdsa生成公钥私钥
//2.签名校验

func main() {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	publicKey := privateKey.PublicKey

	data := "hello world"
	hash := sha256.Sum256([]byte(data))

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		log.Panic()
	}

	//fmt.Printf("r ： %v, len: %d\n",r.Bytes(),len(r.Bytes()))
	//fmt.Printf("s ： %v, len: %d\n",s.Bytes(),len(s.Bytes()))
	//把r, s 进行序列化
	signature := append(r.Bytes(), s.Bytes()...)

	//1.定义两个辅助的big.int
	r1 := big.Int{}
	s1 := big.Int{}
	//2.拆分signature，平均分，前半部分给r, 后半部分给s。(r.Bytes(), s.Bytes() 位数相同)
	r1.SetBytes(signature[:len(signature)/2])
	s1.SetBytes(signature[len(signature)/2:])

	//校验需要三个东西：数据，签名，公钥
	result := ecdsa.Verify(&publicKey, hash[:], &r1, &s1)
	fmt.Println(result)
}