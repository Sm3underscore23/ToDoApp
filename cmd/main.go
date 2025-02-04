package main

import (
	"context"
	"os"
	"os/signal"
	"todoApp/internal/filters"
	"todoApp/internal/handlers"
	"todoApp/pkg/systems"

	"github.com/go-telegram/bot"
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

	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	b.RegisterHandlerMatchFunc(filters.IsHelp, handlers.Help)

	b.Start(ctx)
}
