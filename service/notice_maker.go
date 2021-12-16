package service

type Noticer interface {
	SetText(t string)
	SetRecipientID(r int64)
}

type Notice struct {
	Text        string
	RecipientID int64
}

func NewNotice() *Notice {
	return new(Notice)
} // создание нового извещение

func (n *Notice) SetText(t string) {
	n.Text = t
} // метод для передачи текста сообщения для извещения

func (n *Notice) SetRecipientID(r int64) {
	n.RecipientID = r
} // метод для передачи идентификатора адресата для извещения
