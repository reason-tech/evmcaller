/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var url string
var key string
var chainName string
var chainID int64
var token string
var abiFile string
var valueAmount string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "caller",
	Short: "caller used to call erc721 contract",
	Long:  "caller used to call erc721 contract, not support send value. Hope you enjoy it.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "rpc node url")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "ACCOUNT_PRIVATE_KEY", "private key env name")
	rootCmd.PersistentFlags().StringVarP(&chainName, "chain", "c", "p", "chain name or chain short")
	rootCmd.PersistentFlags().Int64VarP(&chainID, "id", "i", 0, "evm chain id")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174", "interactive token address")
	rootCmd.PersistentFlags().StringVarP(&abiFile, "abi", "a", "abi.abi", "abi file name")
	rootCmd.PersistentFlags().StringVarP(&valueAmount, "value", "v", "0", "send value")

	builtInChains := getBuiltInChains()
	rootCmd.AddCommand(&cobra.Command{
		Use:   "listchain",
		Short: "list support chain.",
		Long:  "List all builtin chains. Use Short or Name in -c",

		Run: func(cmd *cobra.Command, args []string) {
			for _, v := range builtInChains {
				fmt.Println(v.String())
			}
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "getabi",
		Short: "download contract abi.",
		Long:  "download contract abi(require ABI).",
		Run: func(cmd *cobra.Command, args []string) {
			chain, err := getChain(chainName)
			if err != nil {
				fmt.Printf("Not builtin chain: %s, no support getabi\n", chainName)
			}
			url := fmt.Sprintf("%s%s", chain.API, token)
			size := getFile(url, abiFile)
			fmt.Printf("Downloaded a file %s with size %d\n", abiFile, size)
		},
	})

	// TODO:: export abi

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	parsedABI := getABI(ERC721ABI)
	if _, err := os.Stat(abiFile); err == nil {
		fmt.Printf("ABI File exists, parse command from abi.abi\n")
		parsedABI, err = readABI(abiFile)
		if err != nil {
			fmt.Printf("Parsed %s failed, use defaullt ERC721\n", abiFile)
			parsedABI = getABI(ERC721ABI)
		}
	}

	for k, v := range parsedABI.Methods {

		rootCmd.AddCommand(&cobra.Command{
			Use:   k,
			Short: v.Sig,
			Long:  v.String(),

			Run: func(cmd *cobra.Command, args []string) {

				chain, err := getChain(chainName)
				if err != nil {
					fmt.Printf("Not builtin chain, chain id: %d, url: %s\n", chainID, url)
				} else {
					chainID = chain.ChainID
					url = chain.URL
					fmt.Printf("Use chain: %s\n", chain.Name)
				}

				privateKey, err := crypto.HexToECDSA(os.Getenv(key))
				if err != nil {
					log.Fatalln("HexToECDSA: ", err, key)
				}
				address := crypto.PubkeyToAddress(privateKey.PublicKey)
				contract := common.HexToAddress(token)
				value, ok := new(big.Int).SetString(valueAmount, 10)
				if !ok {
					log.Fatal("Value: ", err, valueAmount)
				}
				fmt.Printf("address: %s, contract: %s\n", address, contract)

				client, err := ethclient.Dial(url)
				if err != nil {
					log.Fatalln("Dial: ", err, url)
				}

				var callData []byte
				if len(args) == 0 {
					callData, err = parsedABI.Pack(cmd.Use)
				} else {
					packArgs := getArgs(parsedABI.Methods[cmd.Use], args)
					callData, err = parsedABI.Pack(cmd.Use, packArgs...)
				}
				if err != nil {
					log.Fatalln("Pack: ", err, cmd.Use, args)
				}

				if parsedABI.Methods[cmd.Use].IsConstant() {
					res, err := client.CallContract(context.Background(), ethereum.CallMsg{To: &contract, Value: value, Data: callData}, nil)
					if err != nil {
						log.Fatalln("CallContract: ", err)
					}

					unpack, err := parsedABI.Unpack(cmd.Use, res)
					if err != nil {
						log.Fatalln("Unpack: ", unpack)
					}

					fmt.Println(unpack...)
				} else {
					msg := ethereum.CallMsg{From: address, To: &contract, Value: value, Data: callData}
					gasLimit, err := client.EstimateGas(context.Background(), msg)
					if err != nil {
						log.Fatalln("EstimateGas: ", err)
					}

					gasPrice, err := client.SuggestGasPrice(context.Background())
					if err != nil {
						log.Fatalln("SuggestGasPrice: ", err)
					}

					nonce, err := client.PendingNonceAt(cmd.Context(), address)
					if err != nil {
						log.Fatalln("PendingNonceAt: ", err)
					}

					baseTx := &types.LegacyTx{
						To:       &contract,
						Nonce:    nonce,
						GasPrice: gasPrice,
						Gas:      gasLimit,
						Value:    value,
						Data:     callData,
					}

					signer := types.NewLondonSigner(big.NewInt(chainID))
					signedTx := types.MustSignNewTx(privateKey, signer, baseTx)
					fmt.Println("Broadcasted: ", signedTx.Hash())
					client.SendTransaction(context.Background(), signedTx)
				}
			},
		})
	}
}
