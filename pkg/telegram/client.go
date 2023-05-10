package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	// "path/filepath"
)

var (
	telegramURL = url.URL{Scheme: "https", Host: "api.telegram.org"}
	badRequest  = errors.New("Bad Request")
)

type BotApi struct {
	token    string
	startMsg *StartRequest
}

func initBotApi(token string) *BotApi {
	ba := BotApi{token: token}
	var startMsg StartRequest
	file, _ := os.ReadFile("data/start_msg.json")
	_ = json.Unmarshal(file, &startMsg)
	ba.startMsg = &startMsg
	return &ba
}

func (ba *BotApi) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println(err)
		return nil, badRequest
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	return body, err
}

func (ba *BotApi) registerWebhook(webhookUrl string) bool {
	url := telegramURL.JoinPath("bot"+ba.token, "setWebhook")
	q := url.Query()
	q.Set("url", webhookUrl)
	url.RawQuery = q.Encode()
	_, err := ba.makeRequest(url.String())
	if err != nil {
		return false
	}
	return true
}

func (ba *BotApi) startMessage(chatId int) bool {
	url := telegramURL.JoinPath("bot"+ba.token, "sendMessage")

	values := *ba.startMsg
	values.ID = chatId
	buf, _ := json.Marshal(values)

	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(buf))

	if err != nil {
		log.Println("Error on response.\n", err)
		return false
	}
	if resp.StatusCode != 200 {
		log.Println(badRequest)
		return false
	}
	return true
}
