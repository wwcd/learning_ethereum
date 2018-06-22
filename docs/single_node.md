## 创世配置文件

    wsl@fbi: ~/workdir/geth $ cat genesis.json
    {
      "config": {
            "chainId": 10,
            "homesteadBlock": 0,
            "eip155Block": 0,
            "eip158Block": 0
        },
      "coinbase"   : "0x0000000000000000000000000000000000000000",
      "difficulty" : "0x20000",
      "extraData"  : "",
      "gasLimit"   : "0x2fefd8",
      "nonce"      : "0x0000000000000042",
      "mixhash"    : "0x0000000000000000000000000000000000000000000000000000000000000000",
      "parentHash" : "0x0000000000000000000000000000000000000000000000000000000000000000",
      "timestamp"  : "0x00",
      "alloc"      : {}
    }

## 创建私有链

    wsl@fbi: ~/workdir/geth $ geth init ./genesis.json --datadir "./chain"
    INFO [06-13|10:50:59] Maximum peer count                       ETH=25 LES=0 total=25
    INFO [06-13|10:50:59] Allocated cache and file handles         database=/home/wsl/workdir/geth/chain/geth/chaindata cache=16 handles=16
    INFO [06-13|10:50:59] Writing custom genesis block
    INFO [06-13|10:50:59] Persisted trie from memory database      nodes=0 size=0.00B time=11.433µs gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-13|10:50:59] Successfully wrote genesis state         database=chaindata                                   hash=5e1fc7…d790e0
    INFO [06-13|10:50:59] Allocated cache and file handles         database=/home/wsl/workdir/geth/chain/geth/lightchaindata cache=16 handles=16
    INFO [06-13|10:50:59] Writing custom genesis block
    INFO [06-13|10:50:59] Persisted trie from memory database      nodes=0 size=0.00B time=1.581µs  gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
    INFO [06-13|10:50:59] Successfully wrote genesis state         database=lightchaindata                                   hash=5e1fc7…d790e0

## 启动私有链

    wsl@fbi: ~/workdir/geth $ geth --datadir "./chain" --nodiscover console 2>>eth_output.log
    Welcome to the Geth JavaScript console!

    instance: Geth/v1.8.9-stable/linux-amd64/go1.10.1
     modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

## 查看帐户

    > web3.eth.accounts
    []

## 账户
## 带密码创建

    > web3.personal.newAccount("123456")
    "0x403698a2fc1d1548347167c4a337cab3e987c10c"
    > web3.eth.accounts
    ["0x403698a2fc1d1548347167c4a337cab3e987c10c"]

## 创建

    > web3.personal.newAccount()
    Passphrase:
    Repeat passphrase:
    "0xfe78c1a254ef3758405a501e0a2ca88947bd1700"
    > web3.eth.accounts
    ["0x403698a2fc1d1548347167c4a337cab3e987c10c", "0xfe78c1a254ef3758405a501e0a2ca88947bd1700"]

## 开始挖矿

    > miner.start(1)
    null

挖矿默认使用第一个帐户地址

- `miner.setEtherbase(address)`切换挖矿帐户地址
- `web3.eth.coinbase`查看当前的挖矿帐户地址

## 停止挖矿

    > miner.stop()
    true

## 查看帐户余额

    > web3.eth.getBalance("0x403698a2fc1d1548347167c4a337cab3e987c10c")
    15000000000000000000
    > web3.eth.getBalance("0xfe78c1a254ef3758405a501e0a2ca88947bd1700")
    0

## 帐户别名

    > acc0 = web3.eth.accounts[0]
    "0x403698a2fc1d1548347167c4a337cab3e987c10c"
    > acc1 = web3.eth.accounts[1]
    "0xfe78c1a254ef3758405a501e0a2ca88947bd1700"

## 尝试转帐

    > web3.eth.sendTransaction({from:acc0,to:acc1,value:web3.toWei(3,"ether")})
    Error: authentication needed: password or unlock
        at web3.js:3143:20
        at web3.js:6347:15
        at web3.js:5081:36
        at <anonymous>:1:1

转帐失败是因为涉及到修改区块链的操作都需要用户授权即解锁帐户

## 解锁帐户

    > web3.personal.unlockAccount(acc0,"123456")
    true

## 转帐

    > web3.eth.sendTransaction({from:acc0,to:acc1,value:web3.toWei(3,"ether")})                                                                                                                                                                │·
    "0xd4b55127acc093da23ffb42b7ab6f373ff03217592e942f25c89b73eb7d10057"                                                                                                                                                                       │·

## 查询帐单状态

    > txpool.status
    {
      pending: 1,
      queued: 0
    }

	> txpool.inspect
	{
	  pending: {
		0x403698a2fc1d1548347167c4a337cab3e987c10c: {
		  2: "0xfe78c1a254ef3758405a501e0a2ca88947bd1700: 3000000000000000000 wei + 90000 gas × 18000000000 wei"
		}
	  },
	  queued: {}
	}

查看交易的详细信息

    > web3.eth.getTransaction("0xd4b55127acc093da23ffb42b7ab6f373ff03217592e942f25c89b73eb7d10057")                                                                                                                                            │·
    {                                                                                                                                                                                                                                          │·
      blockHash: "0x0000000000000000000000000000000000000000000000000000000000000000",                                                                                                                                                         │·
      blockNumber: null,                                                                                                                                                                                                                       │·
      from: "0x403698a2fc1d1548347167c4a337cab3e987c10c",                                                                                                                                                                                      │·
      gas: 90000,                                                                                                                                                                                                                              │·
      gasPrice: 18000000000,                                                                                                                                                                                                                   │·
      hash: "0xd4b55127acc093da23ffb42b7ab6f373ff03217592e942f25c89b73eb7d10057",                                                                                                                                                              │·
      input: "0x",                                                                                                                                                                                                                             │·
      nonce: 18,                                                                                                                                                                                                                               │·
      r: "0x7d6a50d6646c3383fe5a82bf6f20a1fe950864536e8563cfe5e6480a6586103f",                                                                                                                                                                 │·
      s: "0x31e34541cb11de2266c2f0faf24546f2cfd6958dc689e53af105ef13070e66e0",                                                                                                                                                                 │·
      to: "0xfe78c1a254ef3758405a501e0a2ca88947bd1700",                                                                                                                                                                                        │·
      transactionIndex: 0,                                                                                                                                                                                                                     │·
      v: "0x37",                                                                                                                                                                                                                               │·
      value: 3000000000000000000                                                                                                                                                                                                               │·
    }

## 查询

	> function checkAllBalances() {
		  var totalBal = 0;
		  for (var acctNum in eth.accounts) {
			  var acct = eth.accounts[acctNum];
			  var acctBal = web3.fromWei(eth.getBalance(acct), "ether");
			  totalBal += parseFloat(acctBal);
			  console.log("  eth.accounts[" + acctNum + "]: \t" + acct + " \tbalance: " + acctBal + " ether");
		  }
		  console.log("  Total balance: " + totalBal + " ether");
	  };
	undefined

	> checkAllBalances()
	  eth.accounts[0]:      0x403698a2fc1d1548347167c4a337cab3e987c10c      balance: 15 ether
	  eth.accounts[1]:      0xfe78c1a254ef3758405a501e0a2ca88947bd1700      balance: 0 ether
	  Total balance: 15 ether
	undefined

## 挖矿触发转帐成功

	> miner.start()
	null
	> checkAllBalances()
	  eth.accounts[0]:      0x403698a2fc1d1548347167c4a337cab3e987c10c      balance: 42 ether
	  eth.accounts[1]:      0xfe78c1a254ef3758405a501e0a2ca88947bd1700      balance: 3 ether
	  Total balance: 45 ether
	undefined
	> miner.stop()
	true
	> checkAllBalances()
	  eth.accounts[0]:      0x403698a2fc1d1548347167c4a337cab3e987c10c      balance: 67 ether
	  eth.accounts[1]:      0xfe78c1a254ef3758405a501e0a2ca88947bd1700      balance: 3 ether
	  Total balance: 70 ether
	undefined

# 查看区块

	> eth.getBlock(0)
	{
	  difficulty: 131072,
	  extraData: "0x",
	  gasLimit: 3141592,
	  gasUsed: 0,
	  hash: "0x5e1fc79cb4ffa4739177b5408045cd5d51c6cf766133f23f7cd72ee1f8d790e0",
	  logsBloom: "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	  miner: "0x0000000000000000000000000000000000000000",
	  mixHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
	  nonce: "0x0000000000000042",
	  number: 0,
	  parentHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
	  receiptsRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
	  sha3Uncles: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
	  size: 507,
	  stateRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
	  timestamp: 0,
	  totalDifficulty: 131072,
	  transactions: [],
	  transactionsRoot: "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
	  uncles: []
	}
