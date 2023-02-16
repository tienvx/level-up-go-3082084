package main

import (
	"fmt"
	"log"
	"sync"
)

// setup constants
const baristaCount = 3
const customerCount = 20
const maxOrderCount = 5

// the total amount of drinks that the bartenders have made
type coffeeShop struct {
	orderCount int
	mu         sync.Mutex

	orderCoffee  chan struct{}
	finishCoffee chan struct{}
	closeShop    chan struct{}
}

// registerOrder ensures that the order made by the baristas is counted
func (p *coffeeShop) registerOrder() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.orderCount++
	if p.orderCount == maxOrderCount {
		close(p.closeShop)
	}
}

// barista is the resource producer of the coffee shop
func (p *coffeeShop) barista(name string) {
	log.Println(name)
	for {
		select {
		case <-p.orderCoffee:
			p.registerOrder()
			log.Printf("%s makes a coffee.\n", name)
			p.finishCoffee <- struct{}{}
		case <-p.closeShop:
			log.Printf("%s stop working.\n", name)
			return
		}
	}
}

// customer is the resource consumer of the coffee shop
func (p *coffeeShop) customer(name string) {
	log.Println(name)
	for {
		select {
		case p.orderCoffee <- struct{}{}:
			log.Printf("%s orders a coffee!", name)
			<-p.finishCoffee
			log.Printf("%s enjoys a coffee!\n", name)
		case <-p.closeShop:
			log.Printf("%s go out of shop.\n", name)
			return
		}
	}
}

func main() {
	log.Println("Welcome to the Level Up Go coffee shop!")
	orderCoffee := make(chan struct{}, baristaCount)
	finishCoffee := make(chan struct{}, baristaCount)
	closeShop := make(chan struct{})
	p := coffeeShop{
		orderCoffee:  orderCoffee,
		finishCoffee: finishCoffee,
		closeShop:    closeShop,
	}
	for i := 0; i < baristaCount; i++ {
		go p.barista(fmt.Sprint("Barista-", i))
	}
	for i := 0; i < customerCount; i++ {
		go p.customer(fmt.Sprint("Customer-", i))
	}
	<-closeShop
	log.Println("The Level Up Go coffee shop has closed! Bye!")
	// WHY ONLY 1 CUSTOMER ORDERS? WHY ONLY 1 BARISTA WORKS?
}
