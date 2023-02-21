package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func main() {
	bill := newBill("car wash")
	fmt.Println("bill: ", bill)
}