package filters

import (
	"todoApp/internal/states"
	"todoApp/internal/texts"

	"github.com/go-telegram/bot/models"
)

func IsStart(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == texts.StartCommand
}

func IsNewEvent(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == texts.NewEventCommand
}

func IsAdd(update *models.Update) bool {
	if update.Message == nil {
		return false
	}
	for key := range texts.NewEventTexts {
		if key == update.Message.Text {
			return true
		}
	}
	return false
}

func CheckState(update *models.Update) bool {
	userID := update.Message.From.ID
	return states.GetState(userID) != ""
}

func IsHelp(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == texts.HelpCommand
}
