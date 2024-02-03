package pubg_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PlayerInfo struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			BanType      string `json:"banType"`
			ClanID       string `json:"clanId"`
			Name         string `json:"name"`
			Stats        interface{} `json:"stats"`
			TitleID      string `json:"titleId"`
			ShardID      string `json:"shardId"`
			PatchVersion string `json:"patchVersion"`
		} `json:"attributes"`
		Relationships struct {
			Matches struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"matches"`
			Assets struct {
				Data []interface{} `json:"data"`
			} `json:"assets"`
		} `json:"relationships"`
		Links struct {
			Self   string `json:"self"`
			Schema string `json:"schema"`
		} `json:"links"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct{} `json:"meta"`
}

type APIService struct {
    apiKey string
}

func CreateAPIService(apiKey string) *APIService {
    return &APIService{apiKey: apiKey}
}

func (a APIService) Tmp() {
    println(a.apiKey)
}

func (a APIService) GetAccountId(nickname string) (string, error){
    // Create a new HTTP request with the appropriate headers
	url := fmt.Sprintf("https://api.pubg.com/shards/kakao/players?filter[playerNames]=%s", nickname)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return "", err
    }

    // Set the required headers
    req.Header.Add("Authorization", "Bearer "+a.apiKey)
    req.Header.Add("Accept", "application/vnd.api+json")

    // Perform the HTTP request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("Error:", err)
        return "", err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return "", fmt.Errorf("error: Unexpected status code %d", resp.StatusCode)
    }
	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}
	
	var playerInfo PlayerInfo
	err2 := json.Unmarshal([]byte(body), &playerInfo)
	if err2 != nil {
		fmt.Println("JSON 언마샬링 오류:", err2)
		return "", err2
	}

	// ID 값 출력
	if len(playerInfo.Data) > 0 {
		playerID := playerInfo.Data[0].ID
		fmt.Println("플레이어 ID:", playerID)
		return playerID, nil
	} 
	return "", nil
}