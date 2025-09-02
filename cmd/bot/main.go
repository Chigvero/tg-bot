package main

import (
	"github.com/Chigvero/tg-bot/internal/service/product"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	producetServce := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		} // Ignore non-message updates

		switch update.Message.Command() {
		case "help":
			printHelp(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, producetServce)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func printHelp(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n",
	)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Here all the products: \n\n")

	products := productService.List()
	for _, p := range products {
		msg.Text += p.Title
		msg.Text += "\n"
	}

	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote:"+inputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
