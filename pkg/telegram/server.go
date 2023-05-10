package telegram

import (
	"fmt"
	"net/http"
	// "sync"
	"encoding/json"
	"log"
	"time"
)

// type database struct {
// 	mapUsers map[int]struct{}
// 	rw     sync.RWMutex
// }

type server struct {
	// db *database
	url    string
	client *BotApi
}

func RunServer(sf *ServerFlags, tf *TelegramFlags, firstRun bool) {
	// database := database{make(map[string]struct{}), sync.RWMutex{}}
	url := fmt.Sprintf("%s:%d", sf.Host, sf.Port)
	botApi := BotApi{token: tf.Token}

	s := server{url: url, client: &botApi}

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.goHandler)
	go func() {
		flag := false
		for !flag {
			time.Sleep(2 * time.Second)
			flag = botApi.registerWebhook(sf.Ngrok)
			log.Println("webhook is run")
		}
	}()
	err := http.ListenAndServe(url, mux)
	if err != nil {
		panic(err)
	}
}

func (s *server) goHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := BotResponse{}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.Msg.Text == "/start" {
		s.client.startMessage(data.Msg.ChatMsg.ID, "Speak Friend and Enter")
	}
}
