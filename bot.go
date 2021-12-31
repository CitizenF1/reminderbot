package main

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	tb "gopkg.in/tucnak/telebot.v2"
)

var botInstance *tb.Bot

func getBotPref() tb.Settings {
	listenPort := ":" + envListenPort

	var poller tb.Poller = &tb.Webhook{
		Listen:   listenPort,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	if botMode != "production" {
		deleteWebhook()
		poller = &tb.LongPoller{Timeout: 1 * time.Second}
	} else {
		if !webhookSet() {
			setWebhook()
		}
	}
	pref := tb.Settings{
		Token:  envBotToken,
		Poller: poller,
	}

	return pref
}

func reminderHelper(m *tb.Message) ([]StoredReminder, bool) {
	done := true
	reminders, err := getUserReminders(m.Sender)
	if err != nil {
		botInstance.Send(m.Chat, "Ошибка поиска напоминаний")
	} else if len(reminders) == 0 {
		botInstance.Send(m.Chat, "Нет активных напоминаний")
	} else {
		done = false
	}
	return reminders, done
}

func forwaredMessage(recipient *tb.User, message *tb.Message) {
	botInstance.Forward(recipient, message)
}

func forwardStoredMessageAfterDelay(id primitive.ObjectID, duration time.Duration) {
	time.Sleep(duration)

	rem, err := getStoredRemindersID(id)
	if err != nil {
		log.Println("Message unable to be retrived")
		return
	}

	message := messageFromStoreddReminder(rem)

	go forwaredMessage(rem.User, &message)
	//removeMessage
	go removeMessageFromDB(id)
}

func forwardMessageAfterDelay(wait Reminder, recipient *tb.User, message *tb.Message) {
	id := storeMessageIntoDB(message, recipient, wait.timestamp)
	forwardStoredMessageAfterDelay(id, wait.duration)
}

func getWaitTime(payload string) (Reminder, error) {
	temp := strings.Join(timeUnits, "|")
	waitExpr := regexp.MustCompile(`(\d+) (` + temp + `)s?`)

	matches := waitExpr.FindStringSubmatch(payload)

	if matches == nil {
		return Reminder{}, errors.New("No matches found")
	}

	quant, _ := strconv.Atoi(matches[1])
	units := matches[2]

	// TODO: Fix how duration is being generated
	seconds := int64(quant * unitMap[units])
	duration := time.Duration(seconds) * time.Second
	timestamp := time.Now().Add(duration)
	return Reminder{units: units, quantity: quant, duration: duration, timestamp: timestamp.Unix()}, nil
}

func confirmReminderSet(wait Reminder, recipient tb.Recipient) {
	stringQuantity := strconv.Itoa(wait.quantity)
	string := "Напоминание установлено для " + stringQuantity + " " + wait.units + "s!"
	botInstance.Send(recipient, string)
}

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
