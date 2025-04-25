package commands

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(message.Chat.ID, "Ketik /menu.")
    bot.Send(msg)
}