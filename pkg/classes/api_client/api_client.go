package api_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
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

type MatchInfo struct {
    Data struct {
        Type       string `json:"type"`
        Id         string `json:"id"`
        Attributes struct {
            PatchVersion string      `json:"patchVersion"`
            BanType      string      `json:"banType"`
            ClanId       string      `json:"clanId"`
            Name         string      `json:"name"`
            Stats        interface{} `json:"stats"`
            TitleId      string      `json:"titleId"`
            ShardId      string      `json:"shardId"`
        } `json:"attributes"`
        Relationships struct {
            Assets struct {
                Data []interface{} `json:"data"`
            } `json:"assets"`
            Matches struct {
                Data []struct {
                    Type string `json:"type"`
                    Id   string `json:"id"`
                } `json:"data"`
            } `json:"matches"`
        } `json:"relationships"`
    } `json:"data"`
    Links struct {
        Self   string `json:"self"`
        Schema string `json:"schema"`
    } `json:"links"`
    Meta interface{} `json:"meta"`
}

type MatchData struct {
    Data struct {
        Type       string      `json:"type"`
        ID         string      `json:"id"`
        Attributes struct {
            Duration     int       `json:"duration"`
            Stats        interface{} `json:"stats"`
            GameMode     string      `json:"gameMode"`
            CreatedAt    time.Time   `json:"createdAt"`
            TitleID      string      `json:"titleId"`
            ShardID      string      `json:"shardId"`
            Tags         interface{} `json:"tags"`
            MapName      string      `json:"mapName"`
            IsCustomMatch bool        `json:"isCustomMatch"`
            MatchType    string      `json:"matchType"`
            SeasonState  string      `json:"seasonState"`
        } `json:"attributes"`
        Relationships struct {
            Rosters struct {
                Data []struct {
                    Type string `json:"type"`
                    ID   string `json:"id"`
                } `json:"data"`
            } `json:"rosters"`
            Assets struct {
                Data []struct {
                    Type string `json:"type"`
                    ID   string `json:"id"`
                } `json:"data"`
            } `json:"assets"`
        } `json:"relationships"`
        Links struct {
            Self   string `json:"self"`
            Schema string `json:"schema"`
        } `json:"links"`
    } `json:"data"`
    Included []struct {
        Type       string `json:"type"`
        ID         string `json:"id"`
        Attributes struct {
            Stats struct {
                DBNOs           int     `json:"DBNOs"`
                Assists         int     `json:"assists"`
                Boosts          int     `json:"boosts"`
                DamageDealt     float64     `json:"damageDealt"`
                DeathType       string  `json:"deathType"`
                HeadshotKills   int     `json:"headshotKills"`
                Heals           int     `json:"heals"`
                KillPlace       int     `json:"killPlace"`
                KillStreaks     int     `json:"killStreaks"`
                Kills           int     `json:"kills"`
                LongestKill     float64     `json:"longestKill"`
                Name            string  `json:"name"`
                PlayerID        string  `json:"playerId"`
                Revives         int     `json:"revives"`
                RideDistance    float64     `json:"rideDistance"`
                RoadKills       int     `json:"roadKills"`
                SwimDistance    float64     `json:"swimDistance"`
                TeamKills       int     `json:"teamKills"`
                TimeSurvived    int     `json:"timeSurvived"`
                VehicleDestroys int     `json:"vehicleDestroys"`
                WalkDistance    float64 `json:"walkDistance"`
                WeaponsAcquired int     `json:"weaponsAcquired"`
                WinPlace        int     `json:"winPlace"`
            } `json:"stats"`
            Actor   string `json:"actor"`
            ShardID string `json:"shardId"`
        } `json:"attributes,omitempty"`
        Relationships struct {
            Team         struct{} `json:"team"`
            Participants struct {
                Data []struct {
                    Type string `json:"type"`
                    ID   string `json:"id"`
                } `json:"data"`
            } `json:"participants"`
        } `json:"relationships,omitempty"`
    } `json:"included"`
    Links struct {
        Self string `json:"self"`
    } `json:"links"`
    Meta interface{} `json:"meta"`
}

type MatchParticipant struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Stats struct {
			DBNOs           int     `json:"DBNOs"`
			Assists         int     `json:"assists"`
			Boosts          int     `json:"boosts"`
			DamageDealt     float64     `json:"damageDealt"`
			DeathType       string  `json:"deathType"`
			HeadshotKills   int     `json:"headshotKills"`
			Heals           int     `json:"heals"`
			KillPlace       int     `json:"killPlace"`
			KillStreaks     int     `json:"killStreaks"`
			Kills           int     `json:"kills"`
			LongestKill     float64     `json:"longestKill"`
			Name            string  `json:"name"`
			PlayerID        string  `json:"playerId"`
			Revives         int     `json:"revives"`
			RideDistance    float64     `json:"rideDistance"`
			RoadKills       int     `json:"roadKills"`
			SwimDistance    float64     `json:"swimDistance"`
			TeamKills       int     `json:"teamKills"`
			TimeSurvived    int     `json:"timeSurvived"`
			VehicleDestroys int     `json:"vehicleDestroys"`
			WalkDistance    float64 `json:"walkDistance"`
			WeaponsAcquired int     `json:"weaponsAcquired"`
			WinPlace        int     `json:"winPlace"`
		} `json:"stats"`
		Actor   string `json:"actor"`
		ShardID string `json:"shardId"`
	} `json:"attributes,omitempty"`
}

type APIService struct {
    ApiKey string
	AccountId string
}

func CreateAPIService(apiKey string) APIService {
    return APIService{ApiKey: apiKey, AccountId: ""}
}

func (a APIService) Tmp() {
    println(a.ApiKey)
}

func (a *APIService) GetAccountId(nickname string) (string, error){
    // Create a new HTTP request with the appropriate headers
	url := fmt.Sprintf("https://api.pubg.com/shards/kakao/players?filter[playerNames]=%s", nickname)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return "", err
    }

    // Set the required headers
    req.Header.Add("Authorization", "Bearer "+a.ApiKey)
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
	a.AccountId = playerInfo.Data[0].ID;
	// ID 값 출력
	if len(playerInfo.Data) > 0 {
		return playerInfo.Data[0].ID, nil
	} 
	return "", nil
}

func (a *APIService) GetLastMatchId() (string, error) {
	if a.AccountId == "" {
		return "", fmt.Errorf("error: AccountId is empty")
	}

	url := fmt.Sprintf("https://api.pubg.com/shards/kakao/players/%s", a.AccountId)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return "", err
    }

    // Set the required headers
    req.Header.Add("Authorization", "Bearer "+a.ApiKey)
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
	
	var matchInfo MatchInfo
	err2 := json.Unmarshal([]byte(body), &matchInfo)
	if err2 != nil {
		fmt.Println("JSON 언마샬링 오류:", err2)
		return "", err2
	}
	var ids []string
	for _, match := range matchInfo.Data.Relationships.Matches.Data {
		ids = append(ids, string(match.Id))
	}
	if len(ids) > 0 {
		return ids[0], nil
	} else {
		return "", nil
	}
}

func (a APIService) GetMatchData(matchId string) ([]MatchParticipant, error) {
	url := fmt.Sprintf("https://api.pubg.com/shards/kakao/matches/%s", matchId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set the required headers
	req.Header.Add("Accept", "application/vnd.api+json")

	// Perform the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: Unexpected status code %d", resp.StatusCode)
	}
	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	
	var matchData MatchData
	err2 := json.Unmarshal([]byte(body), &matchData)
	if err2 != nil {
		fmt.Println("JSON 언마샬링 오류:", err2)
		return nil, err2
	}
	var participants []MatchParticipant
	for _, data := range matchData.Included {
		if data.Type == "participant" {
			if !strings.Contains(data.Attributes.Stats.PlayerID, "ai.") {
				participant := MatchParticipant{
					Type: data.Type,
					ID: data.ID,
					Attributes: data.Attributes,
				}
				participants = append(participants, participant)
			}
			fmt.Println(data.Attributes.Stats.PlayerID)
		}
	}
	fmt.Println(participants)
	return participants, nil
}