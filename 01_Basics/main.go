package main

import (
	"bufio"
	"fmt"
	"hello_world/Ratings"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Do you want to run indefinitely? (yes/no): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "yes" {

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
		} else if input == "no" {
			fmt.Println("Program stopped by user.")
			break
		} else {
			fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		}
	}

}
