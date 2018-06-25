## 代码说明

- `cmd/deploy/main.go`: 部署智能合约
- `cmd/tx/main.go`: 运行智能合约
- `cmd/accessing/main.go`: 查询智能合约
- `mytoken/token.go`: 使用`abigen`生成的智能合约

部署/运行/查询代码中的`key`、`passphrase`、`contractAddress`、`walletAddress`根据实际情况修改

### 运行

    go run cmd/(deploy|tx|accessing)/main.go

## 开发环境搭建

### Windows-64bits

由于`go-ethereum`部分代码依赖c代码，所以需要gcc编译器，windows上需要安装mingw64

#### `go`环境搭建

1. 下载\\10.42.6.66\shared\etherenum\golang\go1.10.1.windows-amd64.zip
2. 解压到C盘根目录
3. 将`C:\go\bin`加入到环境变量`PATH`中
4. 增加环境变量`GOPATH`，内容为`D:\go`
5. 使用`go get gitlab.zte.com.cn/10067372/learning_ethereum`下载代码

#### `mingw64`环境搭建

1. 下载\\10.42.6.66\shared\etherenum\mingw\mingw64.zip
2. 解压到C盘根目录
3. 将`C:\mingw64\bin`加入到环境变量`PATH`中
