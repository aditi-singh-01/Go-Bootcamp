package main

import (
	"fmt"
	"hello_world/Ratings"
)

func main() {
	productRatings := make(map[string]Ratings.Rating)

	var n int
	fmt.Print("How many products do you want to add ratings for? ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("\n========== PRODUCT %d ==========\n", i+1)
		Ratings.InsertProductRatings(productRatings)
	}

	var pid string
	fmt.Print("\nEnter Product ID to view its ratings: ")
	fmt.Scanln(&pid)

	Ratings.PrintProductRatings(productRatings, pid)

}
