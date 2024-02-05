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

func CheckOrCreateUserByAPIKey(apiKey string) (int, error) {
    apiService := utils.CreateAPIService(apiKey)
    // API 키가 유효한지 확인
    _, _err := apiService.GetAccountId("PracticeofEno2")
    if _err != nil {
        return 0, fmt.Errorf("invalid api key")
    }
    // DB에 API 키가 존재하는지 확인
    exist := ExistWithApiKey(apiKey)
    if exist {
        user, err := GetUserByApiKey(apiKey)
        if err != nil {
            return 0, err
        }
        return int(user.ID), nil
    } else {
        //없으면 만듬
        user, err := CreateApiKey(apiKey)
        if err != nil {
            return 0, err
        }
        return int(user.ID), nil
    }
}