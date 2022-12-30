package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestPack(t *testing.T) {
	str := "259"
	i, err := strconv.ParseUint(str, 10, 8)
	fmt.Println(err, i)

	return

	url := "https://polygon-rpc.com"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalln("Dial: ", err, url)
	}

	contract := common.HexToAddress("0xfE8Ea7c29bc64dDa4D03dbFFE4a97b28f830CA07")

	parsedABI := getABI(ERC721ABI)

	interfaceId, err := hex.DecodeString("5b5e139f")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(interfaceId)
	var interfaceId4 [4]byte
	copy(interfaceId4[:], interfaceId[0:4])

	callData, err := parsedABI.Pack("supportsInterface", interfaceId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hex.EncodeToString(interfaceId4[:]))
	fmt.Println(interfaceId4)
	fmt.Println(hex.EncodeToString(callData))

	res, err := client.CallContract(context.Background(), ethereum.CallMsg{To: &contract, Value: big.NewInt(0), Data: callData}, nil)
	if err != nil {
		log.Fatalln("CallContract: ", err)
	}

	unpack, err := parsedABI.Unpack("supportsInterface", res)
	if err != nil {
		log.Fatalln("Unpack: ", unpack)
	}

	fmt.Println(unpack...)

}
