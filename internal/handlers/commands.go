package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todoApp/internal/states"
	"todoApp/internal/texts"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        fmt.Sprintf(texts.WelcomeText, update.Message.From.FirstName),
		ReplyMarkup: texts.MainMenu,
	})
}

func SelectNewEvent(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        fmt.Sprintf(texts.WelcomeText, update.Message.From.FirstName),
		ReplyMarkup: texts.NewEventMenu,
	})
}

func AddTypeEvent(ctx context.Context, b *bot.Bot, update *models.Update) {
	var dataText string
	for key, eventText := range texts.NewEventTexts {
		if key == update.Message.Text {
			states.SetState(update.Message.From.ID, key)
			dataText = eventText
			break
		}
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        dataText,
		ReplyMarkup: texts.NewEventSubmit,
	})
}

func AddEvent(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	state := states.GetState(userID)
	today := time.Now()
	switch state {
	case texts.Today:
		dataText := strings.Split(update.Message.Text, " ")
		if len(dataText) != 2 {
			break
		}
		eventTime := dataText[1]
		eventName := dataText[2]


	case texts.ThisWeek:
	case texts.ThisMouth:
	case texts.ThisYear:
	case texts.Castom:
	}

}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.HelpText,
	})
}

func Default(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update == nil || (update.Message == nil && update.CallbackQuery == nil) {
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.ErrorComandText,
	})
}
