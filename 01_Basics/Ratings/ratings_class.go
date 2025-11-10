package Ratings

import "fmt"

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

func PrintProductRatings(productRatings map[string]Rating, productID string) {
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

func InsertProductRatings(productRatings map[string]Rating) {
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
