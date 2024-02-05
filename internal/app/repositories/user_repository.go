package repositories

import (
	"context"
	"fmt"
	"kill_board/db"
	"kill_board/pkg/utils"
	"math/rand"
)


func ExistWithApiKey(apiKey string) (bool, error) {
    client := utils.GetClient()
    ctx := context.Background()
    _, err := client.User.FindFirst(db.User.APIKey.Equals(apiKey)).Exec(ctx)
    if err != nil {
        fmt.Println(err)
        return false, fmt.Errorf("error: %s", err)
    } else {
        return true, nil
    }
    
}

func GetUserByApiKey(apiKey string) (*db.UserModel, error) {
    client := utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.APIKey.Equals(apiKey)).Exec(ctx)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func CreateApiKey(apiKey string) (*db.UserModel, error) {
    client := utils.GetClient()
    ctx := context.Background()

    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    randomString := make([]byte, 10)
    for i := range randomString {
        randomString[i] = charset[rand.Intn(len(charset))]
    }

    created, err := client.User.CreateOne(
        // required fields
        db.User.APIKey.Set(apiKey),
    ).Exec(ctx)
    if err != nil {
        return nil, fmt.Errorf("error: %s", err)
    }
    return created, nil
}

func DeleteApiKey(apiKey string) (error) {
    client := utils.GetClient()
    ctx := context.Background()
    _, err := client.User.FindUnique(db.User.APIKey.Equals(apiKey)).Delete().Exec(ctx)
    if err != nil {
        return err
    }
    return nil
}