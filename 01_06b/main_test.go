package main

import "testing"

func TestGetBiggestMarketPanic(t *testing.T) {
	defer func() {
		if error := recover(); error == nil {
			panic("method should panic")
		}
	}()
	getBiggestMarket([]User{})
}

func TestGetBiggestMarket(t *testing.T) {
	input := []User{
		{
			Name:    "A",
			Country: "Spain",
		},
		{
			Name:    "B",
			Country: "Germany",
		},
		{
			Name:    "C",
			Country: "France",
		},
		{
			Name:    "D",
			Country: "France",
		},
	}
	country, count := getBiggestMarket(input)
	if country != "France" {
		t.Errorf("Expected %v, got %v", "France", country)
	}
	if count != 2 {
		t.Errorf("Expected %v, got %v", 2, count)
	}
}
