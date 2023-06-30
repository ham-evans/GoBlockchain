package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Block - Defining Block Structure
type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

// CreateBlock - Creating a new block
func CreateBlock (data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}

// Genesis - Creating the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize - BadgerDB interface to serialize block into bytes
func (b *Block) Serialize () []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

// Deserialize - BadgerDB interface to deserialize bytes into block data
func Deserialize (data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)

	return &block
}

// Handle - Gerneral error handler
func Handle (err error) {
	if err != nil {
		fmt.Println("Error!")
		log.Panic(err)
	}
}
