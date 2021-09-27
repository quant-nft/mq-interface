package mqi

type TelegramMessage struct {
	BotName   string  `json:"botName"`
	ChatId    int64   `json:"chatId"`
	Content   string  `json:"content"`
	ParseMode *string `json:"parseMode,omitempty"`
	Preview   *bool   `json:"preview,omitempty"`
}
