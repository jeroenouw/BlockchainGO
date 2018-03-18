# Blockchain with IPFS

## Quick start
Clone this repo: `git clone https://github.com/jeroenouw/BlockchainGO.git`.  
Change directory to this project  
POST your IPFS hash with e.g. Postman: `{"IPFSHash":"[QmQpfRyPKSrDXPQVRkvWDidukWjYs6JKS3P8V6EsEpK7gS]"}`  
Run `go run main.go` to run blockchain.  
Navigate to `http://localhost:8080/`. To see active blockchain.

## Example
```json
[
 {
  "Index": 0,
  "Timestamp": "2018-03-18 16:14:51.012957277 +0100 CET m=+0.005850262",
  "IPFSHash": "",
  "Hash": "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
  "PrevHash": "",
  "Difficulty": 1,
  "Nonce": ""
 },
 {
  "Index": 1,
  "Timestamp": "2018-03-18 16:15:13.911963367 +0100 CET m=+22.904169426",
  "IPFSHash": "QmQpfRyPKSrDXPQVRkvWDidukWjYs6JKS3P8V6EsEpK7gS",
  "Hash": "00e82bcabc8dcd0f9896183434a5b8303b41962319d461f105e4acd6763fb7eb",
  "PrevHash": "5feceb66ffc86f38d952786c6d696c79c2dbc239dd4e91b46729d73a27fb57e9",
  "Difficulty": 1,
  "Nonce": "2"
 }
]
```

## License
MIT