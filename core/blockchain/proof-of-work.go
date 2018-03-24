package blockchain

import (
	"fmt"
	"time"
)

// Generates new block
func generateBlock(oldBlock Block, IPFSHash string) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.IPFSHash = IPFSHash
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	// Proof Of Work
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)

		// Nonce is the hanging value which is added to the concatenated string
		newBlock.Nonce = hex
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), " mining...")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " block succesful")
			newBlock.Hash = calculateHash(newBlock)
			break
		}
	}

	return newBlock, nil
}
