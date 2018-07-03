package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Frontier以太坊V2
	var signer types.FrontierSigner

	// 利用ECDSA(椭圆曲线数字签名算法)生成公私钥
	key, _ := crypto.GenerateKey()
	// 根据公钥生成地址, sha3(publickey)[12:], 取sha3公钥的后20个字节
	addr := crypto.PubkeyToAddress(key.PublicKey)

	// ::::签名过程::::

	// 使用私钥对交易信息进行secp256k1签名,交易信息是下面字段的RLP编码后sha3的hash值
	// tx.data.AccountNonce, tx.data.Price, tx.data.GasLimit, tx.data.Recipient, tx.data.Amount, tx.data.Payload,
	tx, err := types.SignTx(types.NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil), signer, key)
	if err != nil {
		log.Fatalf("can't sign tx: %v\n", err)
	}

	// 打印交易信息,R/S/V为签名信息
	b, err := tx.MarshalJSON()
	if err != nil {
		log.Fatalf("marshal tx failed: %v\n", err)
	}
	log.Println(string(b))

	// ::::检验过程::::

	// 通过签名和交易信息,恢复出公钥,再通过公钥得到地址
	from, err := types.Sender(signer, tx)
	if err != nil {
		log.Fatalf("sender failed: %v\n", err)
	}
	// 检验是否和真实地址是否一致
	if from != addr {
		log.Fatalf("exected from and address to be equal. Got %x want %x", from, addr)
	}
}
