package main

import (
    "log"
    "os"
    "strings"

    "github.com/go-telegram-bot-api/telegram-bot-api"
    "gotesbot/commands"
)

type Config struct {
    BotToken string `json:"botToken"`
    ApiKey   string `json:"apiKey"`
}

func loadConfig() (Config, error) {
    var config Config
    file, err := os.Open("config.json")
    if err != nil {
        return config, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return config, err
}

func main() {
    config, err := loadConfig()
    if err != nil {
        log.Fatal(err)
    }

    bot, err := tgbotapi.NewBotAPI(config.BotToken)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Authorized on account %s", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60

    updates, err := bot.GetUpdatesChan(updateConfig)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        switch {
        case update.Message.Text == "/start":
            commands.Start(bot, update.Message)
        case update.Message.Text == "/ping":
            commands.Ping(bot, update.Message)
        case update.Message.Text == "/menu":
            commands.Menu(bot, update.Message)
        case strings.HasPrefix(update.Message.Text, "/blackbox "):
            query := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/blackbox "))
            commands.Blackbox(bot, update.Message, query, config.ApiKey)
        }
    }
}