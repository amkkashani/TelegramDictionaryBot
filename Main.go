package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	myUsers := make(map[int] int)

	bot, err := tgbotapi.NewBotAPI("1871966618:AAGiXZJ4T_gVgtUIilJxUG2KwnigAmPjq6U")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		fmt.Println(update.Message.From.ID)
		fmt.Println("******")
		if val, ok := myUsers[update.Message.From.ID]; ok {
			myUsers[update.Message.From.ID] = val + 1
		}else{
			myUsers[update.Message.From.ID] = 1
		}
		msg :=tgbotapi.NewMessage(update.Message.Chat.ID, strconv.Itoa(myUsers[update.Message.From.ID]))
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		fmt.Println("***")
		fmt.Println(msg.Text )
		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}