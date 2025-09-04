package commands

import (
	"encoding/json"
	"fmt"
	"github.com/Chigvero/tg-bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	product *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: product,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)

		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)

		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,
			fmt.Sprintf("Parsed Data:%v\n ", parsedData))
		c.bot.Send(msg)
	}

	if update.Message == nil {
		return
	} // Ignore non-message updates

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.ListCommand(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.DefaultBehavior(update.Message)
	}
}

func (c *Commander) DefaultBehavior(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote:"+inputMessage.Text)
	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
