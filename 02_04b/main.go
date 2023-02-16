package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
}

// runAuction and manages the auction for all the items to be sold
// Change the signature of this function as required
func (a *auctioneer) runAuction() {
	for _, item := range items {
		ch := make(chan bid)
		var wg sync.WaitGroup
		wg.Add(len(a.bidders))
		log.Printf("Opening bids for %s!\n", item)
		for _, bidder := range a.bidders {
			go bidder.placeBid(ch, &wg)
		}
		var maxBid bid
		for i := 0; i < len(a.bidders); i++ {
			b := <-ch
			if maxBid.amount < b.amount {
				maxBid = b
			}
		}
		wg.Wait()
		a.bidders[maxBid.bidderID].payBid(maxBid.amount)
	}
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
func (b *bidder) placeBid(ch chan<- bid, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond) // THINK
	amount := rand.Intn(b.wallet)
	ch <- bid{b.id, amount}
	log.Printf("%s place bid %d", b.id, amount)
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
	log.Printf("%s pay bid %d. Remain wallet: %d", b.id, amount, b.wallet)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
		}
		bidders[id] = &b
		//go b.placeBid()
	}
	a := auctioneer{
		bidders: bidders,
	}
	a.runAuction()
	log.Println("The LinkedIn Learning auction has finished!")
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}
