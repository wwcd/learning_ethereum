package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"gitlab.zte.com.cn/10067372/learning_ethereum/dapp/mytoken"
)

const key = `{
  "address": "fe78c1a254ef3758405a501e0a2ca88947bd1700",
  "crypto": {
	"cipher": "aes-128-ctr",
	"ciphertext": "478f7cf0c97926a426f27cbe7be3ef90701075e572139231dd5cbe88531ef23c",
	"cipherparams": {
	  "iv": "51a3e87aa301db4709cbade5bf75e759"
	},
	"kdf": "scrypt",
	"kdfparams": {
	  "dklen": 32,
	  "n": 262144,
	  "p": 1,
	  "r": 8,
	  "salt": "f724c3fee9a123e7dabc21c8b2281b9eaf4b319a1ef618df469143f26e491c6e"
	},
	"mac": "5fc37cd11c695eb4f00a5c5b53f19d43eceaa2cf0590265ba6ce67a4e24c7b16"
  },
  "id": "ca5a614d-241c-43c2-af09-f4ebe34a9924",
  "version": 3
}`

const (
	passphrase = "123456"
	rawurl     = "ws://127.0.0.1:8101"
)

// Deploying contracts to Ethereum
func main() {
	// Create an IPC based RPC connection to a remote node and an authorized transactor
	conn, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := mytoken.DeployToken(auth, conn, big.NewInt(10000), "Contracts in GO!!!")
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	// Don't even wait, check its presence in the local pending state
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}
