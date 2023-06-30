package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

const dbPath = "./tmp/blocks"

// BlockChain - Defining Blockchain Structure
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}
// BlockChainIterator - Struct to Iterate through BlockChain
type BlockChainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}

// AddBlock - Adding a block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)

		lastHash, err = item.ValueCopy(lastHash)

		return err
	})
	Handle(err)

	newBlock := CreateBlock(data, lastHash)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)

		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
}

// InitBlockChain - Initializing the blockchain
func InitBlockChain () *BlockChain {
	var lastHash []byte

	// Diff than tutorial
	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	// Writing first block to DB
	err = db.Update(func(txn *badger.Txn) error { // Closure
		// Is there an existing blockchain? If no, create a new one
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound{ // lh = last hash
			fmt.Println("No existing blockchain found.")
			genesis := Genesis()
			fmt.Println("Genesis proved.")
			err = txn.Set(genesis.Hash, genesis.Serialize()) // Write to DB
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash) // Set last hash to genesis

			lastHash = genesis.Hash

			return err
		} 
		
		// ELSE - if there is an existing blockchain, copy last hash into mem
		item, err := txn.Get([]byte("lh"))
		Handle(err)

		lastHash, err = item.ValueCopy(lastHash[:])

		return err
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}

	return &blockchain
}

// Iterator - Iterate through blockchain, returning each block
func (chain *BlockChain) Iterator () *BlockChainIterator {
	iter := &BlockChainIterator{chain.LastHash, chain.Database}

	return iter
}

// Next - Move to next block in iteration
func (iter *BlockChainIterator) Next () *Block {
	var block *Block
	var encodedBlock []byte

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		Handle(err)

		encodedBlock, err = item.ValueCopy(encodedBlock) 
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
} 