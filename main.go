package main

func main(){
	bc := NewBlockChain()

	cli := CLI{bc}
	cli.Run()

	//bc.AddBlock("A向B转让50枚比特币！")

}
