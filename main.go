package main

import (
	"fmt"
	"goblockchain/blockchain"
	"strconv"
)

func main() {  
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Fist")
	chain.AddBlock("Second")
	chain.AddBlock("Third")

	fmt.Println()
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Block Nonce: %v\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}