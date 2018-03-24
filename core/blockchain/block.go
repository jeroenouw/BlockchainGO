package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

// Mining difficulty (Proof Of Work)
const difficulty = 1

// Block structure
type Block struct {
	Index      int
	Timestamp  string
	IPFSHash   string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

// Message structure
type Message struct {
	IPFSHash string
}

// Prevent data races and generating multiple blocks at same time
var mutex = &sync.Mutex{}

// Hashes block data
func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.IPFSHash + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// Checks if new block is valid
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// Checks if hash is valid
func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

// Handler for writing blocks
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], m.IPFSHash)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	mutex.Unlock()

	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}
