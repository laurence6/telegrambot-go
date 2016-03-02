package utils

import (
	"strconv"

	"github.com/laurence6/telegram-bot-api"
)

func GetMessageChatUserID(message *tgbotapi.Message) string {
	return strconv.Itoa(message.Chat.ID) + "|" + strconv.Itoa(message.From.ID)
}
