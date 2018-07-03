package main

import (
	"bytes"
	"encoding/hex"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	b := &bytes.Buffer{}

	// Yellow Paper
	// Appendix B. Recursive Length Prefix
	// 规则1/2/3针对单个字节串, 规则4/5针对子节串数组, 数组可内嵌数组

	// 规则1: 单个小于128字节,编码为字节本身
	b.Reset()
	rlp.Encode(b, "a")
	log.Println("Rule1 x if ∥x∥ = 1 ∧ x[0] < 128: '0' to", hex.EncodeToString(b.Bytes()))

	// 规则2: 除规则1外,长度小于56的字节串,编码为前缀+字节串本身,前缀为0x80+字节串长度
	b.Reset()
	rlp.Encode(b, "00")
	log.Println("Rule2 (128 + ∥x∥) · x else if ∥x∥ < 56: '00'] to", hex.EncodeToString(b.Bytes()))

	// 规则3: 除规则1/2外,长度大于等于56的字节串,编码为前缀+字节长度+字节串本身,前缀为0xB7+字节长度的字节数
	b.Reset()
	rlp.Encode(b, strings.Repeat("0", 56))
	log.Println("Rule3 (183 + ∥BE(∥x∥)∥) · BE(∥x∥) · x otherwise: '0...56...0' to", hex.EncodeToString(b.Bytes()))

	// 规则4: 数组总长度小于56,编码为前缀+数组成员编码,前缀为0xC0+数组总长度
	b.Reset()
	rlp.Encode(b, []string{"0", "00"})
	log.Println("Rule4 (192 + ∥s(x)∥) · s(x)  if ∥s(x)∥ < 56: ['0', '00'] to", hex.EncodeToString(b.Bytes()))

	// 规则5: 数组总长度大于等于56,编码为前缀+数组总长度+数组成员编码,前缀为0xF8+数组总长度的字节数
	b.Reset()
	rlp.Encode(b, []string{"0", strings.Repeat("0", 56)})
	log.Println("Rule5 (246 + ∥BE(∥s(x)∥)∥) · BE(∥s(x)∥) · s(x)  otherwise: ['0', '0...56...0'] to", hex.EncodeToString(b.Bytes()))

	// ::::模拟对一个交易进行RLP编码::::
	var signer types.FrontierSigner
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	b.Reset()
	tx, err := types.SignTx(types.NewTransaction(0, addr, new(big.Int), 0, new(big.Int), nil), signer, key)
	if err != nil {
		log.Fatalf("can't sign tx: %v\n", err)
	}
	tx.EncodeRLP(b)
	log.Println("Tx:", hex.EncodeToString(b.Bytes()))
}
