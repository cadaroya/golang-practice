package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connecting to ganache network at 8545. Commented out is infura net
	client, err := ethclient.Dial("http://localhost:8545")
	// client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// Initializing Ganache account into a variable
	account := common.HexToAddress("0x74D6691a130984d74ee6342ACFcF192cb0218339")

	// Getting balance of acount in Wei
	balance, err := client.BalanceAt(context.Background(), account, nil)
	fmt.Println(balance)

}
