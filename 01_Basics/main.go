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
	// type ratings struct {
	// 	product string
	// 	rating  int
	// 	comment string
	// }
	// review := ratings{
	// 	product: "Smartphone",
	// 	rating:  5,
	// 	comment: "Outstanding performance and camera quality.",
	// }

	// if review.rating >= 3 {
	// 	fmt.Println("Thank you for your positive feedback!")
	// } else {
	// 	fmt.Println("We appreciate your feedback and will strive to improve.")
	// }

	// fmt.Print("Ratings: ")
	// for i := 0; i < review.rating; i++ {
	// 	fmt.Print("⭐")
	// }
	// fmt.Println()

	//fmt.Printf("Review for %s: Rating %d, Comment: %s\n", review.product, review.rating, review.comment)

	productRatings := make(map[string]Rating)

	productRatings["P1001"] = Rating{
		RatingID: "R001",
		RatingsList: []RatingDetail{
			{UserID: "U001", Rating: 4.5, Comment: "Excellent product!"},
			{UserID: "U002", Rating: 4.0, Comment: "Good quality but a bit pricey."},
			{UserID: "U003", Rating: 5.0, Comment: "Loved it!"},
		},
	}

	productRatings["P2002"] = Rating{
		RatingID: "R002",
		RatingsList: []RatingDetail{
			{UserID: "U010", Rating: 3.8, Comment: "Average quality"},
			{UserID: "U011", Rating: 4.7, Comment: "Highly recommend!"},
		},
	}
	// for pid, rating := range productRatings {
	// 	total := 0.0
	// 	for _, r := range rating.RatingsList {
	// 		total += r.Rating
	// 	}
	// 	rating.AvgRating = total / float64(len(rating.RatingsList))
	// 	productRatings[pid] = rating
	// }

	// printProductRatings(productRatings, "P1001")
	// printProductRatings(productRatings, "P2002")
	// printProductRatings(productRatings, "P9999")
	var n int
	fmt.Print("How many products do you want to add ratings for? ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("\n========== PRODUCT %d ==========\n", i+1)
		insertProductRatings(productRatings)
	}

	// Ask user which product to display
	var pid string
	fmt.Print("\nEnter Product ID to view its ratings: ")
	fmt.Scanln(&pid)

	printProductRatings(productRatings, pid)
}

type RatingDetail struct {
	UserID  string
	Rating  float64
	Comment string
}

type Rating struct {
	RatingID    string
	AvgRating   float64
	RatingsList []RatingDetail
}

func (r Rating) String() string {
	output := fmt.Sprintf("Rating ID: %s | Average Rating: %.2f\n", r.RatingID, r.AvgRating)
	output += "User Reviews:\n"
	for _, review := range r.RatingsList {
		output += fmt.Sprintf("- %v\n", review)
	}
	return output
}

func printProductRatings(productRatings map[string]Rating, productID string) {
	fmt.Printf("\n=== Product %s ===\n", productID)
	if rating, exists := productRatings[productID]; exists {
		fmt.Println(rating)
	} else {
		fmt.Printf("No ratings found for product %s\n", productID)
	}
}

func (r RatingDetail) String() string {
	return fmt.Sprintf("User: %s | Rating: %.1f | Comment: %s", r.UserID, r.Rating, r.Comment)
}

func insertProductRatings(productRatings map[string]Rating) {
	var (
		productID  string
		ratingID   string
		numReviews int
	)

	fmt.Print("Enter Product ID: ")
	fmt.Scanln(&productID)

	fmt.Print("Enter Rating ID: ")
	fmt.Scanln(&ratingID)

	fmt.Print("Enter number of user reviews: ")
	fmt.Scanln(&numReviews)

	rating := Rating{RatingID: ratingID}

	for i := 0; i < numReviews; i++ {
		var userID, comment string
		var score float64

		fmt.Printf("\n--- Review %d ---\n", i+1)
		fmt.Print("Enter User ID: ")
		fmt.Scanln(&userID)
		fmt.Print("Enter Rating (0-5): ")
		fmt.Scanln(&score)
		fmt.Print("Enter Comment: ")
		fmt.Scanln(&comment)

		rating.RatingsList = append(rating.RatingsList, RatingDetail{
			UserID:  userID,
			Rating:  score,
			Comment: comment,
		})
	}

	total := 0.0
	for _, r := range rating.RatingsList {
		total += r.Rating
	}
	rating.AvgRating = total / float64(len(rating.RatingsList))

	productRatings[productID] = rating
	fmt.Println("\n✅ Product rating added successfully!")
}
