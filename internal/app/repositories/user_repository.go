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

func GetUserByRandomString(randomString string) (*db.UserModel, error) {
    client := utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.RandomString.Equals(randomString)).Exec(ctx)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func GetUserByRandomStringWithRelation(randomString string) (*db.UserModel, error) {
    client := utils.GetClient()
    ctx := context.Background()
    user2, err2 := client.User.FindFirst(
        db.User.RandomString.Equals(randomString)).With(db.User.Percents.Fetch()).Exec(ctx)
    if err2 != nil {
        return nil, err2
    } else {
        return user2, nil
    }
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
        db.User.RandomString.Set(string(randomString)),
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

func GetRandomStringByApiKey(apiKey string) (string, error) {
    client := utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.APIKey.Equals(apiKey)).Exec(ctx)
    if err != nil {
        return "", err
    }
    return user.RandomString, nil
}

func ExistRandomString(randomString string) (bool) {
    client := utils.GetClient()
    ctx := context.Background()
    _, err := client.User.FindFirst(db.User.RandomString.Equals(randomString)).Exec(ctx)
    if err != nil {
        return false
    } else {
        return true
    }
}

func UpdateUserCurrentKillByRandomString(randomString string, currentKill int) (error) {
    client := utils.GetClient()
    ctx := context.Background()
    client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
        db.User.CurrentKill.Set(currentKill),
    ).Exec(ctx)
    return nil
}

func UpdateUserTargetKillByRandomString(randomString string, targetKill int) (error) {
    client := utils.GetClient()
    ctx := context.Background()
    client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
        db.User.TargetKill.Set(targetKill),
    ).Exec(ctx)
    return nil
}

func DeletePercentDataByUserID(userID int) (error) {
    client := utils.GetClient()
    ctx := context.Background()
    client.Percent.FindMany(db.Percent.UserID.Equals(userID)).Delete().Exec(ctx)
    return nil
}

func CreatePercentDataByUserID(userId int,  data[] utils.PercentData) (error) {
    client := utils.GetClient()
    ctx := context.Background()
    for _, d := range data {
        _, err := client.Percent.CreateOne(
            db.Percent.Count.Set(d.Count),
            db.Percent.Percent.Set(d.Percent),
            db.Percent.User.Link(db.User.ID.Equals(userId)),
        ).Exec(ctx)
        if err != nil {
            return err
        }
    }
    return nil
}