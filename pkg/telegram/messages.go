package telegram

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

type StartRequest struct {
	ID     int         `json:"chat_id"`
	Text   string      `json:"text"`
	Markup ReplyMarkup `json:"reply_markup"`
}

type ReplyMarkup struct {
	Keyboard [][]Button `json:"keyboard"`
}

type Button struct {
	Text string `json:"text"`
}
