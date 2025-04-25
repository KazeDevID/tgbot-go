package commands

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Ping(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(message.Chat.ID, "hehe")
    bot.Send(msg)
}