package main

import (
	"encoding/hex"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func getArgs(method abi.Method, args []string) []interface{} {
	abiArgs := method.Inputs
	if len(args) != len(abiArgs) {
		log.Fatalf("argument count mismatch: got %d for %d\n", len(args), len(abiArgs))
		// return nil, fmt.Errorf("argument count mismatch: got %d for %d", len(args), len(abiArgs))
	}
	if len(args) == 0 {
		return nil
	}

	s := make([]interface{}, len(args))

	for i := 0; i < len(args); i++ {
		t := abiArgs[i].Type.T
		size := abiArgs[i].Type.Size
		switch t {
		case abi.IntTy:
			s[i] = parseInt(false, size, args[i])
		case abi.UintTy:
			s[i] = parseInt(true, size, args[i])
		case abi.BoolTy:
			boolean, err := strconv.ParseBool(args[i])
			if err != nil {
				log.Fatalln(err)
			}
			s[i] = boolean
		case abi.AddressTy:
			s[i] = common.HexToAddress(args[i])
		case abi.BytesTy:
			s[i] = parsedBytes(false, -1, args[i])
		case abi.FixedBytesTy:
			s[i] = parsedBytes(true, size, args[i])
		}
	}
	return s
}

func parseInt(unsigned bool, size int, input string) interface{} {
	base := 10
	index := 0
	if len(input) >= 2 && (input[:2] == "0x" || input[:2] == "0X") {
		base = 16
		index = 2
	}

	if unsigned {
		switch size {
		case 8:
			val, err := strconv.ParseUint(input[index:], base, size)
			if err != nil {
				log.Fatalln(err)
			}
			return val
		case 16:
			val, err := strconv.ParseUint(input[index:], base, size)
			if err != nil {
				log.Fatalln(err)
			}
			return val
		case 32:
			val, err := strconv.ParseUint(input[index:], base, size)
			if err != nil {
				log.Fatalln(err)
			}
			return val
		case 64:
			val, err := strconv.ParseUint(input[index:], base, size)
			if err != nil {
				log.Fatalln(err)
			}
			return val
		default:
			val, ok := new(big.Int).SetString(input[index:], base)
			if !ok {
				log.Fatalln(ok)
			}
			return val
		}
	}

	switch size {
	case 8:
		val, err := strconv.ParseInt(input[index:], base, size)
		if err != nil {
			log.Fatalln(err)
		}
		return val
	case 16:
		val, err := strconv.ParseInt(input[index:], base, size)
		if err != nil {
			log.Fatalln(err)
		}
		return val
	case 32:
		val, err := strconv.ParseInt(input[index:], base, size)
		if err != nil {
			log.Fatalln(err)
		}
		return val
	case 64:
		val, err := strconv.ParseInt(input[index:], base, size)
		if err != nil {
			log.Fatalln(err)
		}
		return val
	case 256:
		val, ok := new(big.Int).SetString(input[index:], base)
		if !ok {
			log.Fatalln(ok)
		}
		return val
	default:
		log.Fatalln("INVALID INT SIZE", size)
		return nil
	}
}

func parsedBytes(fixed bool, size int, input string) interface{} {
	if len(input) >= 2 && (input[:2] == "0x" || input[:2] == "0X") {
		input = input[2:]
	}

	bytes, err := hex.DecodeString(input)
	if err != nil {
		log.Fatalln(err)
	}

	if !fixed {
		return bytes
	}

	if len(bytes) != size {
		log.Fatalln("INVALID BYTES LENGTH")
	}

	switch size {
	case 1:
		var fixedBytes [1]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 2:
		var fixedBytes [2]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 3:
		var fixedBytes [3]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 4:
		var fixedBytes [4]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 5:
		var fixedBytes [5]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 6:
		var fixedBytes [6]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 7:
		var fixedBytes [7]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 8:
		var fixedBytes [8]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 9:
		var fixedBytes [9]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 10:
		var fixedBytes [10]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 11:
		var fixedBytes [11]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 12:
		var fixedBytes [12]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 13:
		var fixedBytes [13]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 14:
		var fixedBytes [14]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 15:
		var fixedBytes [15]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 16:
		var fixedBytes [16]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 17:
		var fixedBytes [17]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 18:
		var fixedBytes [18]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 19:
		var fixedBytes [19]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 20:
		var fixedBytes [20]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 21:
		var fixedBytes [21]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 22:
		var fixedBytes [22]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 23:
		var fixedBytes [23]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 24:
		var fixedBytes [24]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 25:
		var fixedBytes [25]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 26:
		var fixedBytes [26]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 27:
		var fixedBytes [27]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 28:
		var fixedBytes [28]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 29:
		var fixedBytes [29]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 30:
		var fixedBytes [30]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 31:
		var fixedBytes [31]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	case 32:
		var fixedBytes [32]byte
		copy(fixedBytes[:], bytes)
		return fixedBytes
	default:
		log.Fatalln("INVALID BYTES LENGTH", input)
		return nil
	}
}

func getFile(url, fileName string) int64 {
	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	return size
}
