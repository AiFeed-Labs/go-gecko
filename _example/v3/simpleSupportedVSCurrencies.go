package main

import (
	"fmt"
	"log"

	gecko "github.com/AiFeed-Labs/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	currencies, err := cg.SimpleSupportedVSCurrencies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total currencies", len(*currencies))
	fmt.Println(*currencies)
}
