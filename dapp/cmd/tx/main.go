package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	passphrase      = "123456"
	rawurl          = "/home/wsl/workdir/geth/chain/geth.ipc"
	contractAddress = "0xc6c6d7b4c1b6af38eea8c6c7cc6f08d99f86b920"
	walletAddress   = "0xfe78c1a254ef3758405a501e0a2ca88947bd1700"
)

// Transacting with an Ethereum contract
func main() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	token, err := mytoken.NewToken(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	// Create an authorized transactor and spend 1 unicorn
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	tx, err := token.Transfer(auth, common.HexToAddress(walletAddress), big.NewInt(10000))
	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
}
