package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Value:", rand.Int())

	// const price float32 = 275.00
	// const tax float32 = 27.50
	// const quantity = 2
	// fmt.Println("Total:", quantity*(price+tax))

	// var price2 = price
	// fmt.Println(price2)

	// price3 := 275
	// fmt.Println(price3)

	// price4, tax2, inStock := 275.00, 27.50, true
	// fmt.Println(price4, tax2, inStock)

	first := 100
	var second *int = &first

	first++
	fmt.Println("first:", first)
	fmt.Println("second:", second)
	fmt.Println("*second:", *second)
	*second++
	fmt.Println("*second++:", *second)
}
