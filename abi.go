package main

import (
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func getABI(abiCode string) abi.ABI {
	abi, err := abi.JSON(strings.NewReader(abiCode))
	if err != nil {
		log.Fatal(err)
	}

	return abi
}

func readABI(file string) (abi.ABI, error) {
	abiBytes, err := os.ReadFile(file) // just pass the file name
	if err != nil {
		return abi.ABI{}, err
	}
	abiCode := string(abiBytes)
	parsed, err := abi.JSON(strings.NewReader(abiCode))
	if err != nil {
		return abi.ABI{}, err
	}
	return parsed, nil
}
