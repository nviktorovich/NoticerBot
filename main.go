package main

import (
	"fmt"
	"github.com/nviktorovich/NoticerBot/service"
	"os"
)

// не уверен, что main.go должен быть здесь
func main() {
	bot, err := service.NewTgBot("????")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
