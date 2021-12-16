package demon

import (
	"fmt"
	"github.com/nviktorovich/NoticerBot/service"
	"os"
)

func TelegramNoticer(id int64, text string) {
	notice := service.NewNotice()
	notice.SetRecipientID(id)
	notice.SetText(text)

	bot, err := service.NewTgBot("???")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = bot.Send(*notice); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
