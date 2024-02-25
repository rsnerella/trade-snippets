package main

import (
	"context"
	"log"
	"testing"
)

var kite = &Kite{
	BaseUrl: "https://api.kite.trade",
	Token:   "enctoken pwo1a+V/zxPkxmssyADyp8RCe/Ws0Bl8hVBWZsADFnPt7yaC/1YV2sBjFC4YY+qR8mgkL8gO2q0lmOFlxPDrY4gAngJplM3HEpBvp3JbNz823U+wrCCaOQ==",
}

func TestPlaceOrder(t *testing.T) {
	ctx := context.Background()

	resp, err := kite.PlaceOrder(&ctx, &Order{
		Exchange:                   "NSE",
		TradingSymbol:              "ZOMATO",
		Quantity:                   50,
		MarketProtectionPercentage: 5,
		TickSize:                   0.05,
		TransactionType:            "BUY",
		Product:                    "CNC",
		OrderType:                  "MARKET",
	})
	log.Println(resp, err)

	resp, err = kite.PlaceOrder(&ctx, &Order{
		Exchange:                   "NSE",
		TradingSymbol:              "ZOMATO",
		Price:                      160,
		Quantity:                   50,
		MarketProtectionPercentage: 5,
		TickSize:                   0.05,
		TransactionType:            "BUY",
		Product:                    "CNC",
		OrderType:                  "LIMIT",
	})
	log.Println(resp, err)

	resp, err = kite.PlaceOrder(&ctx, &Order{
		Exchange:                   "NSE",
		TradingSymbol:              "ZOMATO",
		Price:                      160,
		Quantity:                   50,
		MarketProtectionPercentage: 5,
		TickSize:                   0.05,
		TransactionType:            "BUY",
		Product:                    "CNC",
		OrderType:                  "SL",
	})
	log.Println(resp, err)
}
