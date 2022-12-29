

### 前置条件

在命令行导入私钥（本次命令行窗口支持）

```SHELL
export ACCOUNT_PRIVATE_KEY=0000000000000000000000000000000000000000000000000000000000000000
```

或者在命令行下执行，永久支持环境变量

```
cd ~
echo "export ACCOUNT_PRIVATE_KEY=0000000000000000000000000000000000000000000000000000000000000000> .zshrc
source ~/.zshrc
```

```
cd ~
echo "export ACCOUNT_PRIVATE_KEY=0000000000000000000000000000000000000000000000000000000000000000" >> .bashrc
source ~/.bashrc
```

### 使用

#### 查看支持的链路执行

```SHELL
 ./caller listchain
ABI File exists, parse command from abi.abi
Short: m	Name: mumbai	ChainID: 80001	Desc: Polygon Testnet mumbai          	URL: https://rpc-mumbai.maticvigil.com/
Short: e	Name: eth	ChainID: 1	Desc: ETH Mainnet                     	URL: https://mainnet.infura.io/v3/
Short: r	Name: ropsten	ChainID: 5	Desc: ETH Testnet ropsten             	URL: https://goerli.infura.io/v3/
Short: b	Name: bsc	ChainID: 56	Desc: Binance Smart Chain Mainnet     	URL: https://rpc.ankr.com/bsc
Short: bt	Name: bsc-test	ChainID: 97	Desc: Binance Smart Chain Testnet     	URL: https://data-seed-prebsc-1-s1.binance.org:8545
Short: k	Name: kcc	ChainID: 321	Desc: KuCoin Community Chain Mainnet  	URL: https://rpc-mainnet.kcc.network
Short: kt	Name: kcc-test	ChainID: 322	Desc: KuCoin Community Chain Testnet  	URL: https://rpc-testnet.kcc.network
Short: p	Name: polygon	ChainID: 137	Desc: Polygon Mainnet                 	URL: https://polygon-rpc.com
```



#### 自定义命令支持

```SHELL
  -c, --chain string   chain name or chain short (default "p")
  -h, --help           help for caller
  -i, --id int         evm chain id
  -k, --key string     private key env name (default "ACCOUNT_PRIVATE_KEY")
  -t, --token string   interactive token address (default "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
  -u, --url string     rpc node url
```



###### 自定义链路执行（例如 url: 127.0.0.1:8545，chainID: 3） ，输入

> 注意 -p 后面一定要跟一个 listchain 不支持的链路

```SHELL
./caller -p a -u 127.0.0.1:8545 -i 3
```



###### 自定义合约执行

```
./caller -t 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174
```



###### 更改私钥环境变量名称（一般不需要，有自定义需求自用）

```
  -k, --key string     private key env name (default "ACCOUNT_PRIVATE_KEY")
```



#### 执行示例：

1. 调用默认链路 polygon 下的默认合约 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174 的 symbol 方法

```
./caller symbol
Use chain: polygon
address: 0x2a17E171D110E8C62562F84837B1E3b55159b05B, contract: 0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174
USDC
```

2、调用 kcc 链路下指定 token 的 mint 方法

```
./caller -c k -t 0x85412f7ec1112cdda75ec95532b2acb2ca9d5d71 mint 0x2a17E171D110E8C62562F84837B1E3b55159b05B
Use chain: kcc
address: 0x2a17E171D110E8C62562F84837B1E3b55159b05B, contract: 0x85412f7EC1112cdDA75Ec95532B2ACB2ca9D5d71
Broadcasted:  0x9d711db3cf8bbb3813d38071070d051de9bc759cf2d516292eae4121ddb03508
```



#### 自定义 abi 支持

更换支持的 cmd 命令，在同级目录下添加 abi.abi 文件

默认支持 ERC721 的 abi

```
./caller -h
ABI File exists, parse command from abi.abi
caller used to call erc721 contract, not support send value. Hope you enjoy it.

Usage:
  caller [command]

Available Commands:
  DEFAULT_ADMIN_ROLE DEFAULT_ADMIN_ROLE()
  MINTER_ROLE        MINTER_ROLE()
  OPERATOR_ROLE      OPERATOR_ROLE()
  allowance          allowance(address,address)
  approve            approve(address,uint256)
  balanceOf          balanceOf(address)
  blackListAccount   blackListAccount(address)
  blackListed        blackListed(address)
  burn               burn(uint256)
  burnFrom           burnFrom(address,uint256)
  cap                cap()
  completion         Generate the autocompletion script for the specified shell
  decimals           decimals()
  decreaseAllowance  decreaseAllowance(address,uint256)
  getRoleAdmin       getRoleAdmin(bytes32)
  grantRole          grantRole(bytes32,address)
  hasRole            hasRole(bytes32,address)
  help               Help about any command
  increaseAllowance  increaseAllowance(address,uint256)
  listchain          list support chain.
  mint               mint(address,uint256)
  name               name()
  pause              pause()
  paused             paused()
  renounceRole       renounceRole(bytes32,address)
  revokeRole         revokeRole(bytes32,address)
  supportsInterface  supportsInterface(bytes4)
  symbol             symbol()
  totalSupply        totalSupply()
  transfer           transfer(address,uint256)
  transferFrom       transferFrom(address,address,uint256)
  unblackListAccount unblackListAccount(address)
  unpause            unpause()
  withdrawERC20      withdrawERC20(address,address)
  withdrawERC721     withdrawERC721(address,address,uint256)

Flags:
  -c, --chain string   chain name or chain short (default "p")
  -h, --help           help for caller
  -i, --id int         evm chain id
  -k, --key string     private key env name (default "ACCOUNT_PRIVATE_KEY")
  -t, --token string   interactive token address (default "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174")
  -u, --url string     rpc node url

Use "caller [command] --help" for more information about a command.
```



例如 ERC20 的 abi

```
[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"},{"internalType":"uint8","name":"decimals_","type":"uint8"},{"internalType":"uint256","name":"cap_","type":"uint256"},{"internalType":"address","name":"owner_","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"}],"name":"BlackListAccount","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"account","type":"address"}],"name":"Paused","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"role","type":"bytes32"},{"indexed":true,"internalType":"bytes32","name":"previousAdminRole","type":"bytes32"},{"indexed":true,"internalType":"bytes32","name":"newAdminRole","type":"bytes32"}],"name":"RoleAdminChanged","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"role","type":"bytes32"},{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"sender","type":"address"}],"name":"RoleGranted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"bytes32","name":"role","type":"bytes32"},{"indexed":true,"internalType":"address","name":"account","type":"address"},{"indexed":true,"internalType":"address","name":"sender","type":"address"}],"name":"RoleRevoked","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"account","type":"address"}],"name":"UnblackListAccount","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"account","type":"address"}],"name":"Unpaused","type":"event"},{"inputs":[],"name":"DEFAULT_ADMIN_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MINTER_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"OPERATOR_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"blackListAccount","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"blackListed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burnFrom","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"cap","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"}],"name":"getRoleAdmin","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"grantRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"hasRole","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"paused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"renounceRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"revokeRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes4","name":"interfaceId","type":"bytes4"}],"name":"supportsInterface","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"unblackListAccount","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpause","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"address","name":"to","type":"address"}],"name":"withdrawERC20","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"tokenId","type":"uint256"}],"name":"withdrawERC721","outputs":[],"stateMutability":"nonpayable","type":"function"}]
```





