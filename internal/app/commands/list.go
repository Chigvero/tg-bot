package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandData struct {
	Offset int `json:"offset"`
}

func (c *Commander) ListCommand(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Here all the products: \n\n")

	products := c.productService.List()
	for _, p := range products {
		msg.Text += p.Title
		msg.Text += "\n"
	}

	serializedData, _ := json.Marshal(CommandData{Offset: 21})
	keyboad := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("next Page", string(serializedData)),
		),
	)

	msg.ReplyMarkup = keyboad

	c.bot.Send(msg)
}
