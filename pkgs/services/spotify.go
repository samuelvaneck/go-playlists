package services

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const TokenUrl = "https://accounts.spotify.com/api/token"

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetSpotifyToken() (string, error) {
	var token SpotifyToken
	payload := strings.NewReader("grant_type=client_credentials")
	client := &http.Client{}
	req, err := http.NewRequest("POST", TokenUrl, payload)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+string(base64_credentials()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	parseErr := json.Unmarshal([]byte(body), &token)
	if parseErr != nil {
		fmt.Println(parseErr)
		return "", parseErr
	}

	return token.AccessToken, nil
}

func base64_credentials() string {
	credentials := []byte(os.Getenv("SPOTIFY_CLIENT_ID") + ":" + os.Getenv("SPOTIFY_CLIENT_SECRET"))
	return b64.StdEncoding.EncodeToString(credentials)
}

func makeRequest() {
	// ...
}
