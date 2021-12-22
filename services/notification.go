package services

import "sitemon/sender"

type Notificator struct {
	sender sender.Sender
}

func NewNotificator(sender sender.Sender) {
	//..
}
