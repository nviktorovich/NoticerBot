package demon

import "github.com/nviktorovich/NoticerBot/service"

func BotDrive(text string, rID int64) error {

	bot, err := service.NewTgBot("??????")
	if err != nil {
		return err
	}

	notice := service.NewNotice(text, rID)
	if err = bot.Send(*notice); err != nil {
		return err
	}

	return nil
}
