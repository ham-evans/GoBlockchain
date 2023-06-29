package blockchain

// BlockChain - Defining Blockchain Structure
type BlockChain struct {
	Blocks []*Block
}

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

// AddBlock - Adding a block to the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis -Creating the genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain - Initializing the blockchain
func InitBlockChain () *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}