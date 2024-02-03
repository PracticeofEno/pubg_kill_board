package pubg_api

import (
	"kill_board/pubg_api"
	"testing"
)

func TestValidApiKey(t *testing.T) {
	apiKey := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0"
	var apiService = pubg_api.CreateAPIService(apiKey)
	resp, err := apiService.GetAccountId("PracticeofEno2")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Logf("AccounmtId: %s", resp)

}

func TestInValidApiKey(t *testing.T) {
	// expected := 200;
	apiKey := "eyJ0eXAiOiJKV1QiLCJhbGcsiOiJIUzI1NiJ9.eyJqdGkiOiI4ZjY4NjcxMC04ZjIwLTAxM2MtYmQ4Yy0yYWNjZjk1NjI0ZjIiLCJpc3MiOiJnYW1lbG9ja2VyIiwiaWF0IjoxNzA0NTg2ODYxLCJwdWIiOiJibHVlaG9sZSIsInRpdGxlIjoicHViZyIsImFwcCI6IjkxZWU0Y2M4LWY2MTEtNDUxOS05NjQ1LTEzNWJlN2Y4NjkyMiJ9.tadtjH48XZwaRKRSh-ROjHFHuO2dpjrzYASkixz-px0"
	var apiService = pubg_api.CreateAPIService(apiKey)
	_, err := apiService.GetAccountId("PracticeofEno2")
	if err != nil {
		t.Logf("Error: %s", err)
	} else if err == nil {
		t.Errorf("Error: %s", err)
	}

	// t.Logf("Response Status Code : %d", result)
	// if (result == expected) {
	// 	t.Errorf("Expected %d, but got %d", expected, result)
	// }
}