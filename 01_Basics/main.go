package main

import (
	"fmt"
)

func main() {
	// color.Red("Hello, World!")
	// str := "Line1 \n Line 2"
	// color.Blue(str)
	// var comment []string
	//var product, feedback string = "Laptop", "Great performance!"
	// var product string = "Laptop"
	// comment = append(comment, "This is a great product.")
	// comment = append(comment, "Could be improved in battery life.")
	// comment = append(comment, "Excellent value for money.")
	//fmt.Printf("You have provided feedback for %s as: %s\n", product, feedback)
	// 	fmt.Printf("You have provided feedback for %s as: %v\n", product, comment)
	type ratings struct {
		product string
		rating  int
		comment string
	}
	review := ratings{
		product: "Smartphone",
		rating:  5,
		comment: "Outstanding performance and camera quality.",
	}

	if review.rating >= 3 {
		fmt.Println("Thank you for your positive feedback!")
	} else {
		fmt.Println("We appreciate your feedback and will strive to improve.")
	}

	fmt.Print("Ratings: ")
	for i := 0; i < review.rating; i++ {
		fmt.Print("⭐")
	}
	fmt.Println()

	//fmt.Printf("Review for %s: Rating %d, Comment: %s\n", review.product, review.rating, review.comment)
}
