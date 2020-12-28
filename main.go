package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
)

func main() {
	var tokenUrl = "https://crmuat.hanshin.com.tw/CampaignFlowApiService/token/signIn"
	resp, err1 := http.PostForm(tokenUrl, url.Values{"UserId": {"16D1DAB8-1068-EA11-A811-000D3A85426A"}})
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer resp.Body.Close()

	t, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	var token tokenModel

	json.Unmarshal(t, &token)

	fmt.Println(token)
	//my frontend website
	openBrowser("http://localhost/id:e14555d7-9e07-eb11-a813-000d3a851d60^&token:" + token.Data.RefreshToken)
}

func openBrowser(url string) bool {
	var args = []string{"cmd", "/c", "start"}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

type tokenModel struct {
	Token     interface{} `json:"Token"`
	IsSuccess bool        `json:"IsSuccess"`
	Data      struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		UserID       string `json:"UserId"`
	} `json:"Data"`
}
