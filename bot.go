package main

/////////// LEGACY CODE ////////////
// import (
// 	"fmt"
// 	"time"

// 	tb "gopkg.in/tucnak/telebot.v2"
// )

// var (
// 	menu     = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
// 	btnNotif = menu.Text("Тест")
// )

// type Tokens struct {
// 	ApiToken string `json:"api_token"`
// }

// type Reminder struct {
// 	Units     string        `json:"units"`
// 	Duration  time.Duration `json:"duration"`
// 	Timestamp int64         `json:"timestamp"`
// 	Message   string        `json:"message"`
// 	User      *tb.User      `json:"user"`
// }

// type Bot struct {
// 	tbBot  *tb.Bot
// 	tokens *Tokens

// 	Users []*tb.User
// }

// func CreateBot(bot *tb.Bot, tokens *Tokens) *Bot {
// 	return &Bot{
// 		tbBot:  bot,
// 		tokens: tokens,
// 	}
// }

// func (b *Bot) timeChecker() {
// 	for {
// 		location, _ := time.LoadLocation("UTC")
// 		today := time.Now().In(location).Add(time.Hour * 6)
// 		// now := time.Now().Format(time.RFC3339)
// 		fmt.Println(today)

// 		time.Sleep(time.Minute)
// 	}
// }

// func (b *Bot) AuthHandler() {
// 	b.tbBot.Handle("/start", func(m *tb.Message) {
// 		b.tbBot.Send(m.Sender, "Привет"+m.Sender.FirstName, menu)
// 		b.Users = append(b.Users, m.Sender)
// 	})
// }

// func (b *Bot) setReminderHandler() {
// 	b.tbBot.Handle(&btnNotif, func(m *tb.Message) {
// 		b.tbBot.Send(m.Sender, "Message")
// 		b.tbBot.Handle(tb.OnText, func(m *tb.Message) {
// 			if len(m.Text) == 10 {
// 				// str := strings.Split(m.Text, ".")
// 				// date := strings.
// 				// fmt.Println(date)
// 			}
// 		})
// 	})
// }

// func getWaitTime(payload string) (Reminder, error) {
// 	return Reminder{}, nil
// }

// // func (b *Bot) sendUsingUserId(id int, message string) {
// // 	u := &tb.User{
// // 		ID: id,
// // 	}
// // 	b.tbBot.Send(u, message)
// // }

// func (b *Bot) InitMenu() {
// 	menu.Reply(
// 		menu.Row(btnNotif),
// 	)
// }

// func (b *Bot) Init() {
// 	b.InitMenu()
// 	b.AuthHandler()
// 	b.setReminderHandler()
// 	go b.timeChecker()

// 	b.tbBot.Start()
// }