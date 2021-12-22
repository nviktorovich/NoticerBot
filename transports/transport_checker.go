package transports

import (
	"flag"
	"gopkg.in/yaml.v2"
	"os"
)

type Transports struct {
	Clients struct {
		TG   Telegram     `yaml:"telegram"`
		SMS  PhoneMessage `yaml:"sms"`
		Mail Email        `yaml:"mail"`
	} `yaml:"clients"`
}

type Telegram struct {
	Mode   bool   `yaml:"mode"`
	Token  string `yaml:"api_key"`
	ChatID int64  `yaml:"chat_id"`
}

type PhoneMessage struct {
	Mode        bool  `yaml:"mode"`
	PhoneNumber int64 `yaml:"phone_number"`
}

type Email struct {
	Mode    bool   `yaml:"mode"`
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func transportsInits(path string) (*Transports, error) {
	transports := new(Transports)
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &transports); err != nil {
		return nil, err
	}

	return transports, nil
}

func SetConfig() (*Transports, error) {
	var flagPath string
	flag.StringVar(&flagPath, "P", "", "используйте этот флаг, когда нужно загрузить конфигурационный файл не по умолчанию")
	flag.Parse()
	var transports *Transports
	var err error

	if flagPath != "" {
		transports, err = transportsInits(flagPath)

		if err != nil {
			return nil, err
		}

	} else {
		transports, err = transportsInits("./configuration/configuration.yml")
		if err != nil {
			return nil, err
		}
	}
	return transports, nil
}
