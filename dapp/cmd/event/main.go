package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.zte.com.cn/10067372/learning_ethereum/dapp/mytoken"
)

const (
	rawurl          = "ws://127.0.0.1:8101"
	contractAddress = "0xe32effa05a964ebe5c839cb36be9975df45c6819"
)

// Accessing an Ethereum contract
func main() {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	// Instantiate the contract and display its name
	token, err := mytoken.NewToken(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

	sink := make(chan *mytoken.TokenTransfer)

	sub, err := token.WatchTransfer(nil, sink, nil, nil)
	if err != nil {
		log.Fatalf("Failed to watch: %v", err)
	}
	defer sub.Unsubscribe()

	for v := range sink {
		log.Printf("From %s to %s value %v", v.From.String(), v.To.String(), v.Value)
	}
}
