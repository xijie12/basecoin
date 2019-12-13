package main

import (
	"encoding/gob"
	"bytes"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age uint64
}

func main(){
	var p Person
	p.Name = "tom"
	p.Age = 18
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(&p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("编码之后：%v\n", buffer.Bytes())

	var p1 Person
	//dec := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	dec := gob.NewDecoder(&buffer)
	err = dec.Decode(&p1)
	if err != nil {
		log.Fatal("decode:", err)
	}

	fmt.Println(p1)
}