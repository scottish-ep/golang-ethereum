package main

import (
    "os"
	"fmt"
	"log"

    "github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    godotenv.Load()
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_TOKEN"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum network: %v", err)
	}

	contract, err := NewIERC20(common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %v", err)
	}

	amt, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0x0F959508403f8971C811ffa357e306cf43AE57a7"))

    fmt.Println(amt)
	fmt.Println(err)
}
