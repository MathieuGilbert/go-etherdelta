package main

import (
	"log"

	ed "github.com/mathieugilbert/go-etherdelta"
)

func main() {
	service := ed.New(&ed.Options{
		ProviderURI: "wss://mainnet.infura.io/ws",
	})

	orders, err := service.GetOrderBook(&ed.GetOrderBookOpts{
		TokenAddress: "0x0d8775f648430679a709e98d2b0cb6250d2887ef",
	})

	if err != nil {
		panic(err)
	}

	log.Println(orders)
}
