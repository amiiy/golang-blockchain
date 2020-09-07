package main

import (
	"fmt"
	"github.com/amiiy/golang-blockchain/blockchain"
	)


func main() {
	fmt.Println("Initilizeing the block chain ...")
	fmt.Println("getting ready to start")
	fmt.Println("creating the genesis block ...")
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks{
		fmt.Printf("Prev hash is: %x\n",block.PrevHash)
		fmt.Printf("Data in block is: %s\n",block.Data)
		fmt.Printf("hash is: %x\n",block.Hash)

	}
}
