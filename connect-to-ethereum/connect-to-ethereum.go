package main

import (
    "os"
    "context"
    "fmt"
    "log"
    "encoding/json"

    "github.com/joho/godotenv"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    godotenv.Load()
    conn, err := ethclient.Dial("https://mainnet.infura.io/v3/" + os.Getenv("INFURA_TOKEN"))
    if err != nil {
        log.Fatal("Whoops something went wrong!", err)
    }

    ctx := context.Background()
    tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("0x30999361906753dbf60f39b32d3c8fadeb07d2c0f1188a32ba1849daac0385a8"))
    if !pending {
        b, err := json.MarshalIndent(tx, "", "  ")
        if err != nil {
            fmt.Println("error:", err)
        }
        fmt.Printf("%+v", string(b))
    }
}
