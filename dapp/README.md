## 代码说明

- `cmd/deploy/main.go`: 部署智能合约
- `cmd/tx/main.go`: 运行智能合约
- `cmd/accessing/main.go`: 查询智能合约
- `mytoken/token.go`: 使用`abigen`生成的智能合约

部署/运行/查询代码中的`key`、`passphrase`、`contractAddress`、`walletAddress`根据实施情况修改

### 运行

    go run cmd/(deploy|tx|accessing)/main.go
