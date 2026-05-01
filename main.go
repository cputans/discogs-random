package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const discogsBaseURL = "https://api.discogs.com"

func main() {
	username := flag.String("username", "", "Discogs username")
	token := flag.String("token", "", "Discogs personal access token")
	flag.Parse()

	// Get token from environment variable if not provided as flag
	if *token == "" {
		*token = os.Getenv("DISCOGS_TOKEN")
	}

	// Get username from environment variable if not provided as flag
	if *username == "" {
		*username = os.Getenv("DISCOGS_USERNAME")
	}

	if *username == "" || *token == "" {
		fmt.Fprintf(os.Stderr, "Error: username and token are required\n")
		fmt.Fprintf(os.Stderr, "Usage: %s -username <username> -token <token>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Or set DISCOGS_USERNAME and DISCOGS_TOKEN environment variables\n")
		os.Exit(1)
	}

	// Fetch all releases from collection
	releases, err := fetchAllReleases(*username, *token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching releases: %v\n", err)
		os.Exit(1)
	}

	if len(releases) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no releases found in collection\n")
		os.Exit(1)
	}

	// Select random album
	randomRelease := releases[rand.Intn(len(releases))]

	// Display the result
	displayRelease(randomRelease)
}

// fetchAllReleases fetches all releases from the user's collection, handling pagination
func fetchAllReleases(username, token string) ([]CollectionItem, error) {
	var allReleases []CollectionItem
	page := 1

	for {
		url := fmt.Sprintf("%s/users/%s/collection/folders/0/releases?page=%d&per_page=250", discogsBaseURL, username, page)
		resp, err := fetchFromDiscogs(url, token)
		if err != nil {
			return nil, err
		}

		var collResp CollectionResponse
		if err := json.Unmarshal(resp, &collResp); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}

		allReleases = append(allReleases, collResp.Releases...)

		// Check if there are more pages
		if page >= collResp.Pagination.Pages {
			break
		}
		page++
	}

	return allReleases, nil
}

// fetchFromDiscogs makes an authenticated request to the Discogs API
func fetchFromDiscogs(url, token string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers for authentication and user identification
	req.Header.Set("Authorization", fmt.Sprintf("Discogs token=%s", token))
	req.Header.Set("User-Agent", "DiscogRandomAlbumSelector/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// displayRelease prints the random release information
func displayRelease(item CollectionItem) {
	fmt.Println("🎵 Random Album from Your Collection 🎵")
	fmt.Println("=======================================")
	fmt.Printf("Title: %s\n", item.BasicInfo.Title)

	if len(item.BasicInfo.Artists) > 0 {
		fmt.Print("Artists: ")
		for i, artist := range item.BasicInfo.Artists {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(artist.Name)
		}
		fmt.Println()
	}

	if item.BasicInfo.Year > 0 {
		fmt.Printf("Year: %d\n", item.BasicInfo.Year)
	}

	if item.RatingAvg > 0 {
		fmt.Printf("Your Rating: %.1f/5\n", item.RatingAvg)
	}

	fmt.Printf("Added: %s\n", item.DateAdded)
	fmt.Printf("Discogs ID: %d\n", item.ID)
	fmt.Printf("Link: https://www.discogs.com/release/%d\n", item.InstanceID)
}
