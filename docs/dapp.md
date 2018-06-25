
## 智能合约开发步骤

1. 使用[solidity](http://solidity.readthedocs.io/en/develop/)编写智能合约
2. 编译智能合约生成ABI和BIN
3. 部署智能合约
4. 运行智能合约

## 装备工作

安装solidity编译器

    sudo npm install -g solc

    # 使用
    solcjs --bin --abi token.sol

给vscode安装solidity开发插件

安装[go-ethereum](https://github.com/ethereum/go-ethereum), 将编译生成的`build/bin`加入到环境变量`PATH`

	git clone https://github.com/ethereum/go-ethereum
	cd go-ethereum
	make all

    # 使用abigen生成go文件
    abigen --abi token.abi --pkg main --type Token --out token.go --bin token.bin


## 知识点

### BIN

BIN等同lib库，BIN内容会被部署到区块链上，是智能合约的内容

### ABI(Application Binary Interface)

ABI等同lib库的头文件

- itype 方法类型，包括function, constructor, fallback(缺省方法)可以缺省，默认为function
- name 方法名
- inputs 方法参数，它是一个对应数组，数组里的每个对象都是一个参数说明
    * name 参数名
    * type 参数类型
- outputs 方法返回值，格式和inputs类型，如果没有返回值可以缺省
- constant 布尔值，如果为true指明方法不会修改合约的状态变量
- payable 布尔值，标明方法是否可以接收ether

### 智能合约地址

部署完成后，会生成一个智能合约地址，后续运行或查看时，需要使用到智能合约地址

## 样例

样例代码参见[dapp](../dapp)

## 参考

- [Native DApps: Go bindings to Ethereum contracts](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts)
