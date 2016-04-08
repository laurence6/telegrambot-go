package utils

import (
	"strconv"

	"github.com/laurence6/telegram-bot-api"
)

// GetMessageChatUserID retuens "%s|%s", Message.Chat.ID, Message.From.ID .
func GetMessageChatUserID(message *tgbotapi.Message) string {
	return strconv.FormatInt(message.Chat.ID, 10) + "|" + strconv.Itoa(message.From.ID)
}
