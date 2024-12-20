package main

import (
	"fmt"
	"log"

	gecko "github.com/AiFeed-Labs/go-gecko/v3"
	"github.com/AiFeed-Labs/go-gecko/v3/types"
)

func main() {
	cg := gecko.NewClient(nil)
	coin, err := cg.CoinContract(types.CoinNetworkIDEthereum, "0xdac17f958d2ee523a2206206994597c13d831ec7", false, false, false, false, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coin)
}
