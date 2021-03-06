package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramBot struct {
	engine  *tgbotapi.BotAPI
	updChan tgbotapi.UpdatesChannel // канал оставим, если вдруг будет какое-то управление из телеграмма
}

func NewTgBot(token string) (*TelegramBot, error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	updChan := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	return &TelegramBot{
		engine:  bot,
		updChan: updChan,
	}, nil
}

func (bot *TelegramBot) Send(notice Notice) error {
	msg := tgbotapi.NewMessage(notice.RecipientID, notice.Text)
	if _, err := bot.engine.Send(msg); err != nil {
		return err
	}
	return nil
}
