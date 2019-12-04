package main

//4.引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}
//5.定义一个区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}
//定义一个创世块
func GenesisBlock() *Block{
	return NewBlock("Go一期创世块。", []byte{})
}
//6.添加区块
func (bc *BlockChain) AddBlock(data string){
	//获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks) - 1]
	prevHash := lastBlock.Hash

	block := NewBlock(data, prevHash)
	bc.blocks = append(bc.blocks,block)
}
