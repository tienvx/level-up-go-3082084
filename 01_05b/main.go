package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"sort"
)

const path = "items.json"

// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

// matchSales adds the sales procentage of the item
// and sorts the array accordingly.
func matchSales(budget float64, items []SaleItem) []SaleItem {
	match := make([]SaleItem, 0)
	for _, item := range items {
		if item.ReducedPrice <= budget {
			item.SalePercentage = item.ReducedPrice / item.OriginalPrice * 100
			match = append(match, item)
		}
	}
	sort.Slice(match, func(i, j int) bool {
		return match[i].SalePercentage > match[j].SalePercentage
	})
	return match
}

func main() {
	budget := flag.Float64("budget", 0.0,
		"The max budget you want to shop with.")
	flag.Parse()
	items := importData()
	matchedItems := matchSales(*budget, items)
	printItems(matchedItems)
}

// printItems prints the items and their sales.
func printItems(items []SaleItem) {
	log.Println("The BIG sale has started with our amazing offers!")
	if len(items) == 0 {
		log.Println("No items found.:( Try increasing your budget.")
	}
	for i, r := range items {
		log.Printf("[%d]:%s is %.2f %% OFF! Get it now for JUST %.2f!\n",
			i, r.Name, r.SalePercentage, r.ReducedPrice)
	}
}

// importData reads the raffle entries from file and
// creates the entries slice.
func importData() []SaleItem {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
