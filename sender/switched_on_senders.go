package sender

import "sitemon/transports"

func SwitchOnSenders(cfg *transports.Transports) (*Sender, error) {

	var senders []ISender

	if cfg.Clients.TG.Mode {
		tgSender, err := transports.NewTgBot(cfg.Clients.TG.Token)
		if err != nil {
			return nil, err
		}
		senders = append(senders, tgSender)
	}
	//здесь наращиеваем дополнительными сендерами
	s := NewSender(senders)
	return s, nil
}
