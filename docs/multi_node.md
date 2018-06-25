## 节点信息

	NODE1 10.42.6.0
	NODE2 10.42.6.225

## 分别在NODE1和NODE2上启动

	> geth --datadir "./chain" --nodiscover --ipcdisable --port 30301 --ws --wsport 8101 --rpc --rpcport 8102 --networkid 12345 console 2>>eth_output.log

## 查看NODE2节点信息

	> admin.nodeInfo
	{
	  enode: "enode://30020c6118b44f15af9a98db535ceda8f9affa06843cb98cb5b81321524d4a7a67dd108d03b5838b6ddcf595538a8320f77966734d176751fe317196fe6a3b80@[::]:30301",
	  id: "30020c6118b44f15af9a98db535ceda8f9affa06843cb98cb5b81321524d4a7a67dd108d03b5838b6ddcf595538a8320f77966734d176751fe317196fe6a3b80",
	  ip: "::",
	  listenAddr: "[::]:30301",
	  name: "Geth/v1.8.9-stable/linux-amd64/go1.10.1",
	  ports: {
		discovery: 30301,
		listener: 30301
	  },
	  protocols: {
		eth: {
		  config: {
			chainId: 10,
			eip150Hash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			eip155Block: 0,
			eip158Block: 0,
			homesteadBlock: 0
		  },
		  difficulty: 131072,
		  genesis: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
		  head: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
		  network: 12345
		}
	  }
	}

## 新建NODE2用户

	> web3.personal.newAccount("123456")
	"0x8f14c4a5ce50d37353507c8501d37cdf21d75944"

## 在NODE1上增加对端

	> admin.addPeer("enode://30020c6118b44f15af9a98db535ceda8f9affa06843cb98cb5b81321524d4a7a67dd108d03b5838b6ddcf595538a8320f77966734d176751fe317196fe6a3b80@10.42.6.225:30301")
	> admin.peers

