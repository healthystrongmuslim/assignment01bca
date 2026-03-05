package assignment01bca

import (
	"crypto/sha256"
	"time"
)

type block struct {
	prevBlock    *block
	prevHash     [32]byte
	hash         [32]byte // = sha256(prevHash, nonce, merkleRoot)
	nonce        int
	creationTime time.Time
	merkleRoot   [32]byte //the hash of the transaction data (naive i.e. O(n) way for simplicity for now)
	transaction  string
}

func NewBlock(transaction string, nonce int, prev *block) *block {
	nblock := &block{prevBlock: prev, nonce: nonce, creationTime: time.Now(), transaction: transaction}
	if prev != nil {
		nblock.prevHash = prev.hash
	} else {
		nblock.prevHash = [32]byte{0}
	}
	nblock.merkleRoot = sha256.Sum256([]byte(transaction))
	return nblock
}

/*
func ListBlocks() {
	"A method to print all the blocks in a nice format showing block data such as transaction, nonce, previous hash, current block hash"
}

func ChangeBlock() {
	"function to change block transaction of the given block ref"
}
func VerifyChain() {
	"function to verify blockchain in case any changes are made."
}
func CalculateHash(stringToHash string) {
	"function for calculating hash of a block. Hash must be calculated using SHA-256 and must concatenate: transaction + nonce + previousHash + index + timestamp."
}
*/
