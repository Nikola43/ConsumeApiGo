package main

import (
	"encoding/json"
	"fmt"
	. "github.com/logrusorgru/aurora"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
	"time"
)

type Pair struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Tid       int     `json:"tid"`
	Timestamp int     `json:"timestamp"`
}

type Crypto struct {
	Pair [] Pair `json:"ltc_btc"`
}

type ByPrice []Pair
func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPrice) Less(i, j int) bool { return p[i].Price > p[j].Price }



func main() {

	//URL
	var BaseUrl= "https://yobit.net/api/3/trades/ltc_btc?limit=1999"
	var trade Crypto

	//request data
	response, err := http.Get(BaseUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&trade); err != nil {
		log.Println(err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Timestamp\tPrice\tAmount\tTid\tType")

	//sort.Sort(ByPrice(trade.Pair))


	for i := range trade.Pair {
		if trade.Pair[i].Type == "bid" {
			fmt.Fprintf(w, "%s\t%f\t%f\t%d\t%s\n", time.Unix(int64(trade.Pair[i].Timestamp), 0).Format("2006-01-02 15:04:05"), trade.Pair[i].Price, trade.Pair[i].Amount, trade.Pair[i].Tid, Green(trade.Pair[i].Type))
		}
	}
	for i := range trade.Pair {
		if trade.Pair[i].Type == "ask" {
			fmt.Fprintf(w, "%s\t%f\t%f\t%d\t%s\n", time.Unix(int64(trade.Pair[i].Timestamp), 0).Format("2006-01-02 15:04:05"), trade.Pair[i].Price, trade.Pair[i].Amount, trade.Pair[i].Tid, Red(trade.Pair[i].Type))
		}
	}
	//fmt.Println()
	fmt.Fprintln(w)
	w.Flush()
}
