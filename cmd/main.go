package main

import (
	"context"
	"os"
	"os/signal"
	"todoApp/internal/filters"
	"todoApp/internal/handlers"
	"todoApp/pkg/systems"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	token := systems.BotToken()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.Default),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	b.SetMyCommands(ctx, &bot.SetMyCommandsParams{Commands: []models.BotCommand{
		{
			Command:     "/start",
			Description: "Start todoBot",
		},
		{
			Command:     "",
			Description: "Start todoBot",
		},
		{
			Command:     "/start",
			Description: "Start todoBot",
		},
	},
		Scope:        &models.BotCommandScopeDefault{},
		LanguageCode: "en",
	})

	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	b.RegisterHandlerMatchFunc(filters.IsHelp, handlers.Help)
	b.RegisterHandlerMatchFunc(filters.IsNewEvent, handlers.SelectNewEvent)
	b.RegisterHandlerMatchFunc(filters.IsAdd, handlers.AddTypeEvent)
	b.RegisterHandlerMatchFunc(filters.CheckState, handlers.AddEvent)

	b.Start(ctx)
}
