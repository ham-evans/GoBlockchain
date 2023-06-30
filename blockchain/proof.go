package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// Difficulty - Difficulty of POW.
// Could be made variable in future
const Difficulty = 18

// ProofOfWork - Defining the structure for POW
type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

// NewProof - New Proof in POW
func NewProof (b *Block) *ProofOfWork {
	target := big.NewInt(1)

	// Look into LSH
	// https://en.wikipedia.org/wiki/Locality-sensitive_hashing
	target.Lsh(target, uint(256-Difficulty)) 
	

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData - Initializing data with new nonce
func (pow *ProofOfWork) InitData (nonce int) []byte {
	data := bytes.Join([][]byte{pow.Block.PrevHash,  pow.Block.Data, ToHex(int64(nonce), )}, []byte{})
	return data
}

// Run - Running POW algorithm
func (pow *ProofOfWork) Run () (int, []byte) {
	 var intHash big.Int
	 var hash [32]byte

	 nonce := 0

	 for nonce < math.MaxInt64 { // 0 - Infinity (till we find nonce)
		data := pow.InitData(nonce) // Trying with new nonce
		hash = sha256.Sum256(data)

		fmt.Printf("Hash: %x\r", hash[:])

		intHash.SetBytes(hash[:])

		// Until target is bigger than intHash, try a new nonce
		// (aka ensure intHash has enough leading 0's)
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	 }

	 fmt.Println()
	 return nonce, hash[:]
}

// Validate - Validate that the proof given (nonce) is correct
func (pow *ProofOfWork) Validate () bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce) // Create bytes using proposed nonce

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	// Verify that proposed nonce is actually correct
	// This is same check as from Run()
	return intHash.Cmp(pow.Target) == -1 

}

// ToHex - HELPER - converting num to hex
func ToHex (num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	Handle(err)

	return buff.Bytes()
}