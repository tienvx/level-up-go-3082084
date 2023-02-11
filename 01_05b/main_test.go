package main

import (
	"reflect"
	"testing"
)

func TestMatchSales(t *testing.T) {
	budget := 45.
	input := []SaleItem{
		{
			Name:          "50% off",
			OriginalPrice: 100,
			ReducedPrice:  50,
		},
		{
			Name:          "20% off",
			OriginalPrice: 10,
			ReducedPrice:  2,
		},
		{
			Name:          "30% off",
			OriginalPrice: 10,
			ReducedPrice:  3,
		},
	}
	want := []SaleItem{
		{
			Name:           "30% off",
			OriginalPrice:  10,
			ReducedPrice:   3,
			SalePercentage: 30.,
		},
		{
			Name:           "20% off",
			OriginalPrice:  10,
			ReducedPrice:   2,
			SalePercentage: 20.,
		},
	}
	var tests = map[string]struct {
		input []SaleItem
		want  []SaleItem
	}{
		"empty":     {[]SaleItem{}, []SaleItem{}},
		"not empty": {input, want},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := matchSales(budget, test.input)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Expected %v, got %v", test.want, got)
			}
		})
	}
}
