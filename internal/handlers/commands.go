package handlers

import (
	"context"
	"fmt"
	"strconv"
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
		ReplyMarkup: texts.NewEventCencel,
	})
}

func AddEvent(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.From.ID
	state := states.GetState(userID)

	errorSend := func() {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Ошибка формата",
		})
	}

	today := time.Now()

	type patternTime struct {
		year        int
		mounth      int
		day         int
		hour        int
		minute      int
		seconds     int
		miliseconds int
		location    time.Location
	}

	eventTime := patternTime{
		year:        today.Year(),
		mounth:      int(today.Month()),
		day:         today.Day(),
		hour:        today.Hour(),
		minute:      today.Minute(),
		seconds:     0,
		miliseconds: 0,
		location:    *today.Location(),
	}

	userDataLine := update.Message.Text

	if update.Message.Text == texts.CancelComand {
		states.SetState(update.Message.From.ID, "")
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        texts.CancelComand,
			ReplyMarkup: texts.NewEventMenu,
		})
		return
	}

	var userDataSlc []string

	var dataY, datam, dataD, dataH, dataM int

	var err error

	fmt.Println("Start switch 1")
	switch state {
	case texts.Castom:

		fmt.Println("Case Castom")

		userDataSlc = strings.SplitN(userDataLine, " ", 3)

		dataY, err = strconv.Atoi(strings.Split(userDataSlc[0], ".")[2])
		if err != nil {
			errorSend()
			return
		}

		fallthrough

	case texts.ThisYear:

		fmt.Println("Case ThisYear")

		if state != texts.Castom {
			userDataSlc = strings.SplitN(userDataLine, " ", 3)
		}

		datam, err = strconv.Atoi(strings.Split(userDataSlc[0], ".")[1])
		if err != nil {
			errorSend()
			return
		}

		fallthrough

	case texts.ThisMounth, texts.ThisWeek:

		fmt.Println("Case ThisMounth || ThisWeek")

		if state != texts.Castom && state != texts.ThisYear {
			userDataSlc = strings.SplitN(userDataLine, " ", 3)
		}

		dataD, err = strconv.Atoi(userDataSlc[0])
		if err != nil {
			errorSend()
			return
		}

	case texts.Today:

		fmt.Println("Case Today")

		userDataSlc = strings.SplitN(userDataLine, " ", 2)

		dataH, err = strconv.Atoi(strings.Split(userDataSlc[0], ":")[0])
		if err != nil {
			errorSend()
			return
		}

		dataM, err = strconv.Atoi(strings.Split(userDataSlc[0], ":")[1])
		if err != nil {
			errorSend()
			return
		}

	default:
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Ошибка state, напишите в поддержку бота!",
			ReplyMarkup: texts.BackComand,
		})
		return
	}

	eventTitle := userDataSlc[len(userDataSlc)-1]

	chekTime := func(a, b time.Time) bool {
		if a.Hour() < b.Hour() || (a.Hour() == b.Hour() && a.Minute() <= b.Minute()) {
			return false
		} else {
			return true
		}
	}

	fmt.Println("Start switch 2")
	switch state {

	case texts.Castom:

		if dataY < today.Year() {
			errorSend()
			return
		}
		eventTime.year = dataY
		fallthrough

	case texts.ThisYear:

		if state == texts.ThisYear && datam < int(today.Month()) {
			errorSend()
			return
		}
		eventTime.year = dataM
		fallthrough

	case texts.ThisMounth, texts.ThisWeek:

		if state == texts.ThisMounth && dataD < today.Day() {
			fmt.Println("Checker error")
			errorSend()
			return
		}
		eventTime.year = dataD
		fallthrough

	default:

		dataTime, err := time.Parse("15:04", fmt.Sprintf("%d:%d", dataH, dataM))
		if err != nil {
			errorSend()
			return
		}
		if state == texts.Today && chekTime(today, dataTime) {
			errorSend()
			return
		}
		fmt.Println(chekTime(today, dataTime))
		eventTime.hour = dataTime.Hour()
		eventTime.day = dataTime.Day()
		states.SetState(update.Message.From.ID, "")
		fmt.Println(eventTime, eventTitle)
	}

	if states.GetState(userID) != "" {
		errorSend()
		return
	}

	fmt.Println(eventTime, eventTitle)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Событие создано",
		ReplyMarkup: texts.NewEventMenu,
	})
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
		ChatID:      update.Message.Chat.ID,
		Text:        texts.ErrorComandText,
		ReplyMarkup: texts.MainMenu,
	})
}
