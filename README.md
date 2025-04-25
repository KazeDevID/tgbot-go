# Bot Telegram Go

Simple telegram bot using Go language

## Installation

### 1. Condition

Make sure you have Go installed (version 1.15 or later).

### 2. Clone Repository

Clone this repository to your local directory:

```bash
git clone https://github.com/KazeDevID/tgbot-go.git
cd tgbot-go
```

### 3. Install Dependencies

Install required dependencies:

```bash
go mod tidy
```

### 4. Configuration

Create a `config.json` file in your project root directory with the following contents:

```json
{
    "botToken": "YOUR_BOT_TOKEN",
    "apiKey": "YOUR_API_KEY"
}
```
- botToken: you can get your bot token at bot father.
- apikey: You can register and get your apikey at [`here`](https://kaze-apis.my.id/)

## Running the Bot

### 1. In Termux

If you are using Termux, you can run the bot with the following command:

```bash
go run main.go
```

### 2. On VPS

If you are using a VPS, you can do the same steps:

```bash
go run main.go
```

## Note

Make sure your bot has internet access to function properly.
```

### Explanation

- **Installation**: Explains how to install and configure the bot.
- **Running the Bot**: Provides instructions for running the bot on Termux and VPS.
