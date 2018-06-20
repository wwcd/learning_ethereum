package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.zte.com.cn/10067372/learning_ethereum/dapp/mytoken"
)

const (
	rawurl          = "/home/wsl/workdir/geth/chain/geth.ipc"
	contractAddress = "0xc6c6d7b4c1b6af38eea8c6c7cc6f08d99f86b920"
	walletAddress   = "0xfe78c1a254ef3758405a501e0a2ca88947bd1700"
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

	balance, err := token.BalanceOf(nil, common.HexToAddress(walletAddress))
	if err != nil {
		log.Fatalf("query balance error:%v", err)
	}
	fmt.Printf("%s's balance is %s\n", walletAddress, balance)
}
