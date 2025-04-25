package commands

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"

    "github.com/go-telegram-bot-api/telegram-bot-api"
)

const blackboxAPI = "https://kaze-apis.my.id/api/ai/blackbox"

type BlackboxResponse struct {
    Text string `json:"text"`
}

func fetchData(query, apiKey string) (string, error) {
    body, err := json.Marshal(map[string]string{"query": query})
    if err != nil {
        return "", err
    }

    req, err := http.NewRequest("POST", blackboxAPI, bytes.NewBuffer(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("accept", "*/*")
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", err
    }

    var response BlackboxResponse
    body, _ = ioutil.ReadAll(resp.Body)
    if err := json.Unmarshal(body, &response); err != nil {
        return "", err
    }

    return response.Text, nil
}

func Blackbox(bot *tgbotapi.BotAPI, message *tgbotapi.Message, query, apiKey string) {
    if query == "" {
        msg := tgbotapi.NewMessage(message.Chat.ID, "Silakan masukkan query setelah perintah /blackbox.")
        bot.Send(msg)
        return
    }

    responseText, err := fetchData(query, apiKey)
    if err != nil {
        log.Println("Error fetching data:", err)
        responseText = "Error occurred while fetching data."
    }
    msg := tgbotapi.NewMessage(message.Chat.ID, responseText)
    bot.Send(msg)
}