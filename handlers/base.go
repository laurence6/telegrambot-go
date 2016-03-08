package handlers

import "github.com/laurence6/telegram-bot-api"

type HandleFunc func(*tgbotapi.Message) (interface{}, error)
