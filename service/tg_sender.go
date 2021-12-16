package service

type Sender interface {
	Send(notice Notice) error
}
