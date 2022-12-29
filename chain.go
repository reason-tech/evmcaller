package main

import (
	"errors"
	"fmt"
)

type Chain struct {
	Name    string
	Short   string
	Desc    string
	URL     string
	ChainID int64
	API     string
}

var builtInChains = map[string]Chain{
	"bsc":      {Name: "bsc", Short: "b", URL: "https://rpc.ankr.com/bsc", ChainID: 56, Desc: "Binance Smart Chain Mainnet"},
	"bsc-test": {Name: "bsc-test", Short: "bt", URL: "https://data-seed-prebsc-1-s1.binance.org:8545", ChainID: 97, Desc: "Binance Smart Chain Testnet"},
	"kcc":      {Name: "kcc", Short: "k", URL: "https://rpc-mainnet.kcc.network", ChainID: 321, Desc: "KuCoin Community Chain Mainnet"},
	"kcc-test": {Name: "kcc-test", Short: "kt", URL: "https://rpc-testnet.kcc.network", ChainID: 322, Desc: "KuCoin Community Chain Testnet"},
	"polygon":  {Name: "polygon", Short: "p", URL: "https://polygon-rpc.com", ChainID: 137, Desc: "Polygon Mainnet", API: "https://api.polygonscan.com/api?module=contract&action=getabi&address="},
	"mumbai":   {Name: "mumbai", Short: "m", URL: "https://rpc-mumbai.maticvigil.com/", ChainID: 80001, Desc: "Polygon Testnet mumbai"},
	"eth":      {Name: "eth", Short: "e", URL: "https://mainnet.infura.io/v3/", ChainID: 1, Desc: "ETH Mainnet", API: "https://api.etherscan.io/api?module=contract&action=getabi&address="},
	"ropsten":  {Name: "ropsten", Short: "r", URL: "https://goerli.infura.io/v3/", ChainID: 5, Desc: "ETH Testnet ropsten"},
}

var shortMap = map[string]string{
	"b":  "bsc",
	"bt": "bsc-test",
	"k":  "kcc",
	"kt": "kcc-test",
	"p":  "polygon",
	"m":  "mumbai",
	"e":  "eth",
	"r":  "ropsten",
}

func (c *Chain) String() string {
	return fmt.Sprintf("Short: %s\tName: %s\tChainID: %d\tDesc: %-32s\tURL: %s", c.Short, c.Name, c.ChainID, c.Desc, c.URL)
}

func getBuiltInChains() map[string]Chain {
	return builtInChains
}

func getChain(nameOrShort string) (Chain, error) {

	if len(nameOrShort) <= 2 {
		if name, ok := shortMap[nameOrShort]; ok {
			return builtInChains[name], nil
		}
	}

	if chain, ok := builtInChains[nameOrShort]; ok {
		return chain, nil
	}

	return Chain{}, errors.New("NOT BUILTIN CHAIN")
}
