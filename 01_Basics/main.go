package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	fmt.Print("Do you want to continue? (yes/no): ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(strings.ToLower(input))

	// 	if input == "yes" {

	// 		productRatings := make(map[string]Ratings.Rating)

	// 		var n int
	// 		fmt.Print("How many products do you want to add ratings for? ")
	// 		fmt.Scanln(&n)

	// 		for i := 0; i < n; i++ {
	// 			fmt.Printf("\n========== PRODUCT %d ==========\n", i+1)
	// 			Ratings.InsertProductRatings(productRatings)
	// 		}

	// 		var pid string
	// 		fmt.Print("\nEnter Product ID to view its ratings: ")
	// 		fmt.Scanln(&pid)

	// 		Ratings.PrintProductRatings(productRatings, pid)
	// 	} else if input == "no" {
	// 		fmt.Println("Program stopped by user.")
	// 		break
	// 	} else {
	// 		fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
	// 	}
	// }

	urls := []string{
		"https://gdg.community.dev/gdg-cochin/",
		"https://golang.org",
		"https://httpstat.us/500",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.twitter.com/",
		"https://www.instagram.com/",
		"https://site-not-present.io",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://www.github.com/",
		"https://www.stackoverflow.com/",
		"https://www.reddit.com/",
	}
	// var wg sync.WaitGroup

	// for _, url := range urls {
	// 	wg.Add(1)
	// 	go verifyUrls(url, &wg)
	// }
	// fmt.Println("\nAll the sites are checked!")
	// wg.Wait()

	rootCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	results := make(chan Result, len(urls)) // channel for worker -> printer
	var workersWG sync.WaitGroup            // wait group to wait for all checkers

	// Start the printer goroutine.
	var printerWG sync.WaitGroup
	printerWG.Add(1)
	go printer(rootCtx, results, &printerWG)

	// Spawn one goroutine per URL to check it.
	for _, u := range urls {
		workersWG.Add(1)
		go func(url string) {
			defer workersWG.Done()
			status, dur, err := check(rootCtx, url) // go check()
			results <- Result{URL: url, Status: status, Duration: dur, Err: err}
		}(u)
	}

	// Wait for all workers, then close results so printer can finish.
	workersWG.Wait()
	close(results)
	printerWG.Wait()

	fmt.Println("All done.")

}

// Result holds the outcome of a single check.
type Result struct {
	URL      string
	Status   string // "UP" or "DOWN"
	Duration time.Duration
	Err      error
}

// check performs an HTTP GET with a per-request timeout and returns "UP" or "DOWN".
func check(ctx context.Context, url string) (string, time.Duration, error) {

	// 	defer wg.Done()
	// response, err := http.Get(url)
	// if err != nil {
	// 	fmt.Printf("X %s is DOWN\n", url)
	// 	return
	// }
	// defer response.Body.Close()
	// if response.StatusCode == 200 {
	// 	fmt.Printf("ss is UP\n", url)
	// } else {
	// 	fmt.Printf("X%s is DOWN\n", url)
	// }
	// Per-request timeout (adjust as needed).
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		return "DOWN", 0, err
	}

	client := &http.Client{
		// A small transport and overall timeout keep things responsive.
		Timeout: 6 * time.Second,
	}

	start := time.Now()
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		return "DOWN", elapsed, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return "UP", elapsed, nil
	}
	return "DOWN", elapsed, fmt.Errorf("status code %d", resp.StatusCode)
}

// printer drains the results channel and prints human-friendly lines.
func printer(ctx context.Context, in <-chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Context canceled (e.g., Ctrl+C) — stop printing.
			return
		case r, ok := <-in:
			if !ok {
				// Channel closed: nothing more to print.
				return
			}
			if r.Err != nil {
				fmt.Printf("[%-4s] %-30s (%v) err=%v\n", r.Status, r.URL, r.Duration, r.Err)
			} else {
				fmt.Printf("[%-4s] %-30s (%v)\n", r.Status, r.URL, r.Duration)
			}
		}
	}
}

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// func checkSite(url string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("❌ %s is DOWN\n", url)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode == 200 {
// 		fmt.Printf("✅ %s is UP\n", url)
// 	} else {
// 		fmt.Printf("❌ %s is DOWN\n", url)
// 	}
// }

// func main() {
// 	urls := []string{
// 		"https://gdg.community.dev/gdg-cochin/",
// 		"https://golang.org",
// 		"https://httpstat.us/500",
// 		"https://www.google.com/",
// 		"https://www.facebook.com/",
// 		"https://www.twitter.com/",
// 		"https://www.instagram.com/",
// 		"https://site-not-present.io",
// 		"https://www.youtube.com/",
// 		"https://www.linkedin.com/",
// 		"https://www.github.com/",
// 		"https://www.stackoverflow.com/",
// 		"https://www.reddit.com/",
// 	}

// 	// Create a context with 10 second timeout
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()

// 	//wg is a Synchronization primitive that allows you to wait for a collection of goroutines to finish executing.
// 	var wg sync.WaitGroup

// 	for _, url := range urls {
// 		wg.Add(1)

// 		go checkSite(url, &wg)
// 	}

// 	// Channel to signal when all goroutines are done
// 	done := make(chan struct{})
// 	go func() {
// 		wg.Wait()
// 		close(done)
// 	}()

// 	// Wait for either all checks to complete or timeout
// 	select {
// 	case <-done:
// 		fmt.Println("\nAll checks completed!")
// 	case <-ctx.Done():
// 		fmt.Println("\n⏱️  Timeout reached! Some checks may still be running...")
// 	}
// }
