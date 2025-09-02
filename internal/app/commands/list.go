package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) ListCommand(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Here all the products: \n\n")

	products := c.productService.List()
	for _, p := range products {
		msg.Text += p.Title
		msg.Text += "\n"
	}

	c.bot.Send(msg)
}
