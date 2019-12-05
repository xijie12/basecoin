package main

import (
	"strings"
	"fmt"
)

func main(){
	str1 := []string{"1","2","3"}
	res := strings.Join(str1,"+")
	fmt.Printf("res: %s\n",res)
}
