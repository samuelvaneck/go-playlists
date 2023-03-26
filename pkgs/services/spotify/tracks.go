package services

import "fmt"

const SearchEndpoint = "https://api.spotify.com/v1/search?q="

func SearchTrack() {
	//
}

func SearchUrl(searchParams string) (string, error) {
	if searchParams == "" {
		return "", fmt.Errorf("searchParams is empty")
	}

	return fmt.Sprintf("%s%s&type=track", SearchEndpoint, searchParams), nil
}
