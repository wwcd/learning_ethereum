// https://github.com/ethereum/wiki/wiki/Patricia-Tree

package main

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

func main() {
	t, _ := trie.New(common.Hash{}, trie.NewDatabase(ethdb.NewMemDatabase()))

	// 为了让压缩字典树的槽位,对KEY进行了十六进制编码,eg 0x10编码为[0x01, 0x00, 0x10(结束符)]
	// 	shortNode: 仅有一个子节点的节点
	// 	fullNode: 最多有16子节点的节点,槽位17为本节点的数据节点
	// 	valueNode: 承载数据的节点
	// 	hashNode: 特殊节点

	t.Update([]byte{0x10}, []byte("foo"))
	/*
		ShortNode{key:0x010010}
					 |
					 V
				ValueNode("foo")
	*/

	t.Update([]byte{0x10, 0x10}, []byte("bar"))
	/*
		ShortNode{key:0x0100}
					   |
					   V
					FullNode[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, a, b, c, d, e, f, G]
								|											 |
								V											 V
						ShortNode{key:0010, val}		   				ValueNode("foo")
											 |
											 V
										ValueNode("bar")

	*/

	t.Update([]byte{0x10, 0x80}, []byte("boom"))
	/*
		ShortNode{key:0x010010}
					   |
					   V
					FullNode[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, a, b, c, d, e, f, G]
								|				     |						 |
								V				     V						 V
						ShortNode{key:0010}  ShortNode{key:0010} 		ValueNode("foo")
									 |      		   	   |
									 V   				   V
								ValueNode("bar")		ValueNode("boom")

	*/

	log.Println(t.Hash().String())

	iter := trie.NewIterator(t.NodeIterator(nil))
	for iter.Next() {
		log.Printf("key=%s value=%s\n", hex.EncodeToString(iter.Key), string(iter.Value))
	}
}
