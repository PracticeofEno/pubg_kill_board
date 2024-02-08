package service

import (
	"fmt"
	"kill_board/db"
	"kill_board/internal/app/repositories"
	"kill_board/pkg/utils"
)

func ExistWithApiKey(apiKey string) (bool) {
    exist, _ := repositories.ExistWithApiKey(apiKey)
    return exist
}

func GetUserByApiKey(apiKey string) (*db.UserModel, error) {
    user, err := repositories.GetUserByApiKey(apiKey)
    if err != nil {
        fmt.Printf("GetUserByApiKey erorr : %s", err)
        return nil, err
    }
    return user, nil
}

func GetUserByRandomStringWithRelation(randomString string) (*db.UserModel, error) {
    user, err := repositories.GetUserByRandomStringWithRelation(randomString)
    if err != nil {
        fmt.Printf("GetUserByRandomString error : %s", err)
        return nil, err
    }
    return user, nil
}

func CreateApiKey(apiKey string) (*db.UserModel, error) {
    user, err := repositories.CreateApiKey(apiKey)
    if err != nil {
        fmt.Printf("CreateApiKey error : %s", err)
        return nil, err
    }
    return user, nil
}

func DeleteApiKey(apiKey string) (error) {
    err := repositories.DeleteApiKey(apiKey)
    if err != nil {
        fmt.Printf("DeleteApiKey error : %s", err)
        return err
    }
    return nil
}

func CheckOrCreateUserByAPIKey(apiKey string) (string, error) {
    apiService := utils.CreateAPIService(apiKey)
    // API 키가 유효한지 확인
    _, _err := apiService.GetAccountId("PracticeofEno2")
    if _err != nil {
        return "", fmt.Errorf("invalid api key")
    }
    // DB에 API 키가 존재하는지 확인
    exist := ExistWithApiKey(apiKey)
    if exist {
        user, err := GetUserByApiKey(apiKey)
        if err != nil {
            return "", err
        }
        return user.RandomString, nil
    } else {
        //없으면 만듬
        user, err := CreateApiKey(apiKey)
        if err != nil {
            return "", err
        }
        return user.RandomString, nil
    }
}

func GetRandomString(apiKey string) (string, error) {
    randomString, err := repositories.GetRandomStringByApiKey(apiKey)
    if err != nil {
        fmt.Printf("GetRandomString error : %s", err)
        return "", err
    }
    return randomString, nil
}

func ExistRandomString(randomString string) (bool) {
    return repositories.ExistRandomString(randomString)
}

func RecreatePercentData(data utils.PercentReqeust) {
    user, err := repositories.GetUserByRandomString(data.RandomString)
    if err != nil {
        fmt.Printf("error : %s", err)
        return
    }
    err2 := repositories.CreatePercentDataByUserID(user.ID, data.PercentData)
    if err != nil {
        fmt.Printf("error : %s", err2)
        return
    }
}

func UpdateUserData(data utils.UpdateUserData) {
    user, err := repositories.GetUserByRandomString(data.RandomString)
    if err != nil {
        fmt.Printf("error : %s", err)
        return
    }
    repositories.DeletePercentDataByUserID(user.ID)
    repositories.CreatePercentDataByUserID(user.ID, data.PercentData)
    repositories.UpdateUserCurrentKillByRandomString(data.RandomString, data.CurrentKill)
    repositories.UpdateUserTargetKillByRandomString(data.RandomString, data.TargetKill)
}