package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/ethdb"
)

func main() {
	// 加载本地的LevelDB
	db, err := ethdb.NewLDBDatabase("/home/wsl/workdir/geth/chain/geth/chaindata", 768, 1024)
	if err != nil {
		log.Fatalf("new ldb failed: %v\n", err)
	}

	// 构建区块链
	chainConfig, _, err := core.SetupGenesisBlock(db, nil)
	if err != nil {
		log.Fatalf("failed to get chain config: %v\n", err)
	}
	blockchain, err := core.NewBlockChain(db, nil, chainConfig, ethash.NewFaker(), vm.Config{})
	if err != nil {
		log.Fatalf("failed to create new chain manager: %v\n", err)
	}

	// 获取一个指定的区块
	block := blockchain.GetBlockByNumber(3397)
	if block == nil {
		log.Fatalf("failed to get block")
	}
	log.Printf("root of block is %s\n", block.Root().String())

	// 对一个区块进行txhash检验, block_validator.go
	txs := block.Transactions()
	txHash := types.DeriveSha(txs)
	log.Printf("txRoot of block is %s, validate is %s\n", block.TxHash().String(), txHash.String())

	// 获取区块链的stateDB
	stateDB, err := blockchain.StateAt(blockchain.CurrentBlock().Root())
	if err != nil {
		log.Fatalf("failed to get state db: %v\n", err)
	}

	// 从stateDB中获取指定地址的余额
	addr := "0xfe78C1A254eF3758405A501e0a2cA88947BD1700"
	blance := stateDB.GetBalance(common.HexToAddress(addr))
	log.Println(addr, blance)

	// 获取所有个人帐户交易次数及余额
	dump := stateDB.RawDump()

	for k, v := range dump.Accounts {
		if v.Code == "" {
			log.Println(k, v.Nonce, v.Balance)
		}
	}

	// 获取所有智能合约帐户
	// for k, v := range dump.Accounts {
	// 	if v.Code != "" {
	// 		fmt.Println(k, v.Code)
	// 	}
	// }
}
