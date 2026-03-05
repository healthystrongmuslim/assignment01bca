package assignment01bca

import (
	"crypto/sha256"
	"strconv"
	"time"
)

type block struct {
	prevBlock    *block
	prevHash     string
	hash         string // = sha256(prevHash, nonce, merkleRoot, creationTime, index)
	nonce        int
	index        int
	creationTime time.Time
	merkleRoot   string //the hash of the transaction data (naive i.e. O(n) way for simplicity for now)
	transaction  string
}

func byte32toStr(data [32]byte) string {
	return string(data[:])
}

func CalculateHash(stringToHash string) string {
	return byte32toStr(sha256.Sum256([]byte(stringToHash)))
}

func NewBlock(transaction string, nonce int, prev *block, index int) *block {
	nblock := &block{prevBlock: prev, nonce: nonce, creationTime: time.Now(), transaction: transaction, index: index}
	if prev != nil {
		nblock.prevHash = prev.hash
	} else {
		nblock.prevHash = ""
	}
	nblock.merkleRoot = CalculateHash(transaction)
	header := nblock.prevHash + strconv.Itoa(nonce) + nblock.merkleRoot + nblock.creationTime.String() + strconv.Itoa(index)
	nblock.hash = CalculateHash(header)
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
*/
