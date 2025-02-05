package texts

import (
	"github.com/go-telegram/bot/models"
)

const (
	StartCommand    = "/start"
	BackComand      = "Назад"
	ListCommand     = "Список событий"
	NewEventCommand = "Новое событие"
	HelpCommand     = "Помощь"

	Today     = "Сегодня"
	ThisWeek  = "На неделе"
	ThisMouth = "В этом месяце"
	ThisYear  = "В этом году"
	Castom    = "Выбрать дату"

	WelcomeText     = "Привет, %s.\nTodoApp - это прилежение для планирования дел.\nЕсли ты не знаком с функционалом нажми кнопку \"Помощь\"."
	HelpText        = "Все команды TodoApp:\n\"Список событий\" - отображения списка событий/дел\n\"Новое событие\" - добавление события"
	ErrorComandText = "Неизвестная команда!"
)

var NewEventTexts = map[string]string{
	Today:     "Введите событие на сегодня:\n<ЧАСЫ:МИНУТЫ> <НАЗВАНИЕ СОБЫТИЯ>\n\nПример: 15:30 Встреча с командой",
	ThisWeek:  "Введите событие на эту неделю:\n<ДЕНЬ НЕДЕЛИ> <ЧАСЫ:МИНУТЫ> <НАЗВАНИЕ СОБЫТИЯ>\n\nПример: Понедельник 10:00 Планерка",
	ThisMouth: "Введите событие в этом месяце:\n<ДЕНЬ> <ЧАСЫ:МИНУТЫ> <НАЗВАНИЕ СОБЫТИЯ>\n\nПример: 25 18:00 День рождения",
	ThisYear:  "Введите событие в этом году:\n<ДЕНЬ.МЕСЯЦ> <ЧАСЫ:МИНУТЫ> <НАЗВАНИЕ СОБЫТИЯ>\n\nПример: 31.12 23:50 Новый год",
	Castom:    "Введите событие с выбранной датой:\n<ДЕНЬ.МЕСЯЦ.ГОД> <ЧАСЫ:МИНУТЫ> <НАЗВАНИЕ СОБЫТИЯ>\n\nПример: 01.09.2024 08:00 Первое сентября",
}

var MainMenu = &models.ReplyKeyboardMarkup{
	Keyboard: [][]models.KeyboardButton{
		{{Text: ListCommand}},
		{{Text: NewEventCommand}},
		{{Text: HelpCommand}},
	},
	ResizeKeyboard: true,
}

var NewEventMenu = &models.ReplyKeyboardMarkup{
	Keyboard: [][]models.KeyboardButton{
		{{Text: Today}, {Text: ThisWeek}},
		{{Text: ThisMouth}, {Text: ThisYear}},
		{{Text: Castom}},
		{{Text: BackComand}},
	},
	ResizeKeyboard: true,
}

var NewEventSubmit = &models.ReplyKeyboardMarkup{
	Keyboard: [][]models.KeyboardButton{
		{{Text: "Подтвердить"}},
	},
	ResizeKeyboard:  true,
	OneTimeKeyboard: true,
}
