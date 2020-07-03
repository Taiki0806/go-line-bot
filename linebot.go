package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

func main() {
	handler, err := httphandler.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					// log確認用
					fmt.Println(message.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						fmt.Println("textでエラー発生")
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf("スタンプ嬉しい！ありがとう！！")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						fmt.Println("stampでエラー発生")
						log.Print(err)
					}
				}
			}
		}
	})
	http.Handle("/callback", handler)

	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
