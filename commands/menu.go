package commands

import (
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Menu(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(message.Chat.ID, "Menu:\n/start - Memulai bot\n/ping - Mengecek status bot\n/blackbox <query> - Menggunakan fitur blackbox")
    bot.Send(msg)
}