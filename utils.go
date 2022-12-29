package main

import (
	"encoding/hex"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func getArgs(method abi.Method, args []string) []interface{} {

	if len(args) == 0 {
		return nil
	}

	s := make([]interface{}, len(args))

	for i := 0; i < len(args); i++ {
		t := method.Inputs[i].Type.T
		size := method.Inputs[i].Type.Size
		switch t {
		case abi.IntTy:
			switch size {
			case 8:
				val, err := strconv.ParseInt(args[i], 10, 8)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = int8(val)
			case 16:
				val, err := strconv.ParseInt(args[i], 10, 16)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = int16(val)
			case 32:
				val, err := strconv.ParseInt(args[i], 10, 32)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = int32(val)
			case 64:
				val, err := strconv.ParseInt(args[i], 10, 64)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = int64(val)
			default:
				var ok bool
				s[i], ok = new(big.Int).SetString(args[i], 10)
				if !ok {
					log.Fatalln(ok)
				}
			}
		case abi.UintTy:
			switch size {
			case 8:
				val, err := strconv.ParseUint(args[i], 10, 8)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = uint8(val)
			case 16:
				val, err := strconv.ParseUint(args[i], 10, 16)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = uint16(val)
			case 32:
				val, err := strconv.ParseUint(args[i], 10, 32)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = uint32(val)
			case 64:
				val, err := strconv.ParseUint(args[i], 10, 64)
				if err != nil {
					log.Fatalln(err)
				}
				s[i] = uint64(val)
			default:
				big, ok := new(big.Int).SetString(args[i], 10)
				if !ok {
					log.Fatalln(ok)
				}
				s[i] = big
			}
		case abi.BoolTy:
			boolean, err := strconv.ParseBool(args[i])
			if err != nil {
				log.Fatalln(err)
			}
			s[i] = boolean
		case abi.AddressTy:
			s[i] = common.HexToAddress(args[i])
		case abi.BytesTy:
			arr, err := hex.DecodeString(args[i])
			if err != nil {
				log.Fatalln(err)
			}
			s[i] = arr

		case abi.FixedBytesTy:
			arr, err := hex.DecodeString(args[i])
			if err != nil {
				log.Fatalln(err)
			}
			if len(arr) != size {
				log.Fatalln("INVALID BYTES LENGTH")
			}

			fixedArr := reflect.ArrayOf(size, reflect.TypeOf(byte(0)))
			reflect.Copy(reflect.ValueOf(arr), reflect.ValueOf(fixedArr))
			s[i] = reflect.ValueOf(fixedArr).Interface
		}
	}
	return s
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
