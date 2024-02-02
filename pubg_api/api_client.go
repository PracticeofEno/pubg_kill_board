package pubg_api

import (
	"fmt"
	"io"
	"net/http"
)

type APIService struct {
	apiKey string
}

func CreateAPIService(apiKey string) *APIService {
    return &APIService{apiKey: apiKey}
}

func (a APIService) Tmp() {
	println(a.apiKey)
}

func (a APIService) GetPlayer() {
	// Create a new HTTP request with the appropriate headers
	req, err := http.NewRequest("GET", "https://api.pubg.com/shards/kakao/players?filter[playerNames]=PracticeofEno2", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the required headers
	req.Header.Add("Authorization", "Bearer "+a.apiKey)
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	if (resp.StatusCode != 200) {
		fmt.Println("Error:", resp.Status)
		return
	}
	fmt.Println(string(body))
}