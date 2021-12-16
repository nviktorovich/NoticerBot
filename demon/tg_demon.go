package demon

import "github.com/nviktorovich/NoticerBot/service"

func BotDrive(bot *service.TelegramBot, text string, rID int64) error {

	notice := service.NewNotice(text, rID)
	if err := bot.Send(*notice); err != nil {
		return err
	}
	return nil
}