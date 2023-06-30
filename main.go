package main

import (
	"flag"
	"fmt"
	"goblockchain/blockchain"
	"os"
	"runtime"
	"strconv"
)

// CommandLine - Building CLI for command line
type CommandLine struct {
	blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage () {
	fmt.Println("Usage:")
	fmt.Println(" add - block BLOCK_DATA - add this to the blockchain")
	fmt.Println(" print - Prints the blocks in the chain")
}

func (cli *CommandLine) validateArgs () {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock (data string) {
	cli.blockchain.AddBlock(data)
	fmt.Println("Added Block.")
	fmt.Println()
}

func (cli *CommandLine) printChain () {
	iter := cli.blockchain.Iterator()

	fmt.Println()
	for{
		block := iter.Next()
		
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		// fmt.Printf("Block Nonce: %v\n", block.Nonce)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 { // We have hit genesis block
			break
		}
	}
}

func (cli *CommandLine) run () {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add": 
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print": 
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {  
	chain := blockchain.InitBlockChain()
	defer chain.Database.Close() // Properly close DB before main ends

	cli := CommandLine{chain}
	cli.run()




}