package handlers

import (
	"context"
	"todoApp/internal/texts"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.Welcome + bot.EscapeMarkdown(update.Message.From.FirstName) + "!",
	})
}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	menu := &models.ReplyKeyboardMarkup{
		Keyboard: 
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        texts.Help,
	})
}

func Default(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.ErrorComand,
	})
}
