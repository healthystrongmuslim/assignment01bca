package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	PrevBlock    *Block
	PrevHash     string
	Hash         string // = sha256(prevHash, nonce, merkleRoot, creationTime, index)
	Nonce        int
	Index        int
	CreationTime time.Time
	MerkleRoot   string //the hash of the transaction data (naive i.e. O(n) way for simplicity for now)
	Transaction  string
}

func byte32toStr(data [32]byte) string {
	return hex.EncodeToString(data[:])
}

func CalculateHash(stringToHash string) string {
	return byte32toStr(sha256.Sum256([]byte(stringToHash)))
}

func NewBlock(transaction string, nonce int, prev *Block, index int) *Block {
	nblock := &Block{PrevBlock: prev, Nonce: nonce, CreationTime: time.Now(), Transaction: transaction, Index: index}
	if prev != nil {
		nblock.PrevHash = prev.Hash
	} else {
		nblock.PrevHash = ""
	}
	nblock.MerkleRoot = CalculateHash(transaction)
	header := nblock.PrevHash + strconv.Itoa(nonce) + nblock.MerkleRoot + strconv.FormatInt(nblock.CreationTime.Unix(), 10) + strconv.Itoa(index)
	nblock.Hash = CalculateHash(header)
	return nblock
}
