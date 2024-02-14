package auto_kill

import (
	"fmt"
	"kill_board/internal/app/repositories"
	"kill_board/pkg/classes/api_client"
	"kill_board/pkg/utils/hub"
	"strings"
	"time"
)

type Worker struct {
    ApiService      api_client.APIService
    Nickname        string
    RandomString   string
    LastMatchId     string
    PrevLastMatchId string
}

func NewWorker(apikey string, nickname string, randomString string) Worker {
    apiService := api_client.CreateAPIService(apikey)
    user, err := repositories.GetUserByApiKey(apikey)
    if err != nil {
        panic(err)
    }
    apiService.GetAccountId(user.Nickname)
    lastMatchId, err := apiService.GetLastMatchId()
    if err != nil {
        panic(err)
    }
    return Worker{
        ApiService:      apiService,
        Nickname:        nickname,
        LastMatchId:     lastMatchId,
        PrevLastMatchId: lastMatchId,
        RandomString:  randomString,
    }
}

func (w *Worker) Run() {
    var count int = 0;
	var end_count int = 720;
    defer func() {
        fmt.Printf("%s - 종료합니다\n", w.Nickname)
		repositories.ChangeActiveByRandomString(w.RandomString, false)
    }()
        
    for {
        if count > end_count {
            fmt.Printf("%s - 1시간 동안 전적이 검색되지 않았습니다. 종료합니다 \n", w.Nickname)
            break;
        }

        lastMatchId, err := w.ApiService.GetLastMatchId()
        if err != nil {
            fmt.Printf("Error - GetLastMatchId: %s\n", err)
            time.Sleep(10 * time.Second)
            count += 1
            continue
        }
        if strings.EqualFold(w.LastMatchId, lastMatchId) {
            fmt.Printf("%s - No new match\n", w.Nickname)
            time.Sleep(10 * time.Second)
            count += 1
            continue
        }
        
        participants, err := w.ApiService.GetMatchData(lastMatchId)
        if err != nil {
            fmt.Printf("Error - GetMatchData: %s\n", err)
            time.Sleep(10 * time.Second)
            count += 1
            continue
        }
        for _, participant := range participants {
            if strings.EqualFold(participant.Attributes.Stats.Name, w.Nickname) {
                repositories.AddUserCurrentKillByApiKey(w.ApiService.ApiKey, participant.Attributes.Stats.Kills)
                fmt.Printf("New kill: %d\n", participant.Attributes.Stats.Kills)
            }
        }
        w.LastMatchId = lastMatchId
        hub_tmp := hub.GetHub()
        hub_tmp.GetWsClient(w.RandomString).Emit("pong2", "haha")
        count = 0
    }
}