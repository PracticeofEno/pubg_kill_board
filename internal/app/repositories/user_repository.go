package repositories

import (
	"context"
	"fmt"
	"kill_board/db"
	"kill_board/db_utils"
	"kill_board/pkg/utils/dto"
	"math/rand"
)

func ChangeActiveByRandomString(randomString string, active bool) (error) {
	client := db_utils.GetClient()
	ctx := context.Background()
	client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
		db.User.Active.Set(active),
	).Exec(ctx)
	return nil
}

func ExistWithApiKey(apiKey string) (bool, error) {
    client := db_utils.GetClient()
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
    client := db_utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.APIKey.Equals(apiKey)).Exec(ctx)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func GetUserByRandomString(randomString string) (*db.UserModel, error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.RandomString.Equals(randomString)).Exec(ctx)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func GetUserByRandomStringWithRelation(randomString string) (*db.UserModel, error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    user2, err2 := client.User.FindFirst(
        db.User.RandomString.Equals(randomString)).With(
			db.User.Percents.Fetch().OrderBy(db.Percent.Percent.Order(db.ASC)),
		).Exec(ctx)
    if err2 != nil {
        return nil, err2
    } else {
        return user2, nil
    }
}

func CreateApiKey(apiKey string) (*db.UserModel, error) {
    client := db_utils.GetClient()
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
    client := db_utils.GetClient()
    ctx := context.Background()
    _, err := client.User.FindUnique(db.User.APIKey.Equals(apiKey)).Delete().Exec(ctx)
    if err != nil {
        return err
    }
    return nil
}

func GetRandomStringByApiKey(apiKey string) (string, error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    user, err := client.User.FindFirst(db.User.APIKey.Equals(apiKey)).Exec(ctx)
    if err != nil {
        return "", err
    }
    return user.RandomString, nil
}

func ExistRandomString(randomString string) (bool) {
    client := db_utils.GetClient()
    ctx := context.Background()
    _, err := client.User.FindFirst(db.User.RandomString.Equals(randomString)).Exec(ctx)
    if err != nil {
        return false
    } else {
        return true
    }
}

func AddUserCurrentKillByApiKey(apiKey string, addKill int) (error) {
	client := db_utils.GetClient()
	ctx := context.Background()
	client.User.FindUnique(db.User.APIKey.Equals(apiKey)).Update(
		db.User.CurrentKill.Increment(addKill),
	).Exec(ctx)
	return nil
}

func AddUserTargetKillByApiKey(apiKey string, addKill int) (error) {
	client := db_utils.GetClient()
	ctx := context.Background()
	client.User.FindUnique(db.User.APIKey.Equals(apiKey)).Update(
		db.User.TargetKill.Increment(addKill),
	).Exec(ctx)
	return nil
}

func AddUserCurrentKillByRandomString(randomString string, addKill int) (error) {
	client := db_utils.GetClient()
	ctx := context.Background()
	client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
		db.User.CurrentKill.Increment(addKill),
	).Exec(ctx)
	return nil
}

func AddUserTargetKillByRandomString(randomString string, addKill int) (error) {
	client := db_utils.GetClient()
	ctx := context.Background()
	client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
		db.User.TargetKill.Increment(addKill),
	).Exec(ctx)
	return nil
}

func UpdateUserCurrentKillByRandomString(randomString string, currentKill int) (error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
        db.User.CurrentKill.Set(currentKill),
    ).Exec(ctx)
    return nil
}

func UpdateUserTargetKillByRandomString(randomString string, targetKill int) (error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    client.User.FindUnique(db.User.RandomString.Equals(randomString)).Update(
        db.User.TargetKill.Set(targetKill),
    ).Exec(ctx)
    return nil
}

func DeletePercentDataByUserID(userID int) (error) {
    client := db_utils.GetClient()
    ctx := context.Background()
    client.Percent.FindMany(db.Percent.UserID.Equals(userID)).Delete().Exec(ctx)
    return nil
}

func CreatePercentDataByUserID(userId int,  data[] dto.PercentData) (error) {
    client := db_utils.GetClient()
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