package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	telegramURL = url.URL{Scheme: "https", Host: "api.telegram.org"}
	badRequest  = errors.New("Bad Request")
)

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Message struct {
	ID      int    `json:"message_id"`
	ChatMsg Chat   `json:"chat"`
	FromMSG From   `json:"from"`
	Date    int    `json:"date"`
	Text    string `json:"text"`
}

type BotResponse struct {
	ID  int     `json:"update_id"`
	Msg Message `json:"message"`
}

type BotApi struct {
	token string
}

func (ba *BotApi) makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n", err)
		return nil, err
	}
	if resp.StatusCode == 404 {
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

func (ba *BotApi) startMessage(chatId int, text string) bool {
	url := telegramURL.JoinPath("bot"+ba.token, "sendMessage")

	values := map[string]any{"chat_id": chatId, "text": text}
	buf, _ := json.Marshal(values)
	resp, err := http.Post(url.String(), "application/json", bytes.NewBuffer(buf))

	if err != nil {
		log.Println("Error on response.\n", err)
		return false
	}
	if resp.StatusCode == 404 {
		log.Println(badRequest)
		return false
	}

	return true
}
