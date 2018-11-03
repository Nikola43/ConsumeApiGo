package main

type CryptoTrade struct {
	Date          int     `json:"date"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	Tid           int     `json:"tid"`
	PriceCurrency string  `json:"price_currency"`
	Item          string  `json:"item"`
	TradeType     string  `json:"trade_type"`
}