package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Nonce        int
	Transaction  string
	PreviousHash string
	Hash         string
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(nonce int, transaction string, previousHash string) *Block {
	block := &Block{
		Nonce:        nonce,
		Transaction:  transaction,
		PreviousHash: previousHash,
		Hash:         CalculateHash(fmt.Sprintf("%s%d%s", transaction, nonce, previousHash)),
	}
	return block
}

func ChangeBlock(block *Block, newTransaction string) {
	block.Transaction = newTransaction
	block.Hash = CalculateHash(fmt.Sprintf("%d%s%s", block.Nonce, block.Transaction, block.PreviousHash))
}
func DisplayBlocks(bc *Blockchain) {
	for _, block := range bc.Blocks {
		fmt.Printf("Nonce: %d, Transaction: %s, Previous Hash: %s, Current Hash: %s\n", block.Nonce, block.Transaction, block.PreviousHash, block.Hash)
	}
}

func VerifyChain(bc *Blockchain) bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != CalculateHash(fmt.Sprintf("%s%d%s", currentBlock.Transaction, currentBlock.Nonce, currentBlock.PreviousHash)) {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hashInBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashInBytes[:])
}

func main() {
	// Create a new blockchain
	blockchain := &Blockchain{
		Blocks: []*Block{
			{
				Nonce:        0,
				Transaction:  "Florance",
				PreviousHash: "",
				Hash:         "",
			},
		},
	}

	// Adding 5 dummy blocks
	for i := 0; i < 5; i++ {
		transaction := fmt.Sprintf("user%d to user%d", i, i+1)
		nonce := i
		previousHash := blockchain.Blocks[len(blockchain.Blocks)-1].Hash
		newBlock := NewBlock(nonce, transaction, previousHash)
		blockchain.Blocks = append(blockchain.Blocks, newBlock)
	}

	// Display all blocks
	DisplayBlocks(blockchain)

	// Verify the blockchain
	isValid1 := VerifyChain(blockchain)
	if isValid1 {
		fmt.Println(" valid Block-chain ")
	} else {
		fmt.Println("invalid Block-chain .")
	}

	fmt.Printf("\n\n\n")

	// Change the transaction of the Second  block
	ChangeBlock(blockchain.Blocks[2], "user2 to user5")

	// Display all blocks
	DisplayBlocks(blockchain)
	// Verify the blockchain
	isValid2 := VerifyChain(blockchain)
	if isValid2 {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
