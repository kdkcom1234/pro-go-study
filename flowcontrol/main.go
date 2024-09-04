package main

import "fmt"

func main() {

	kayakPrice := 275.00
	fmt.Println("Price:", kayakPrice)

	if kayakPrice > 100 {
		fmt.Println("Price is greater than 100")
	}

	product := "kayak"
	for index, character := range product {
		fmt.Println("index", index, "character", string(character))
	}

	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("index", index, "element", element)
	}
}
