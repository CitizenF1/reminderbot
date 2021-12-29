package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	menu      = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
	btnRemind = menu.Text("Напомни")
	btnList   = menu.Text("Список")
)

type Reminder struct {
	units     string
	quantity  int
	duration  time.Duration
	timestamp int64
}

type StoredReminder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    int64              `bson:"chat_id"`
	MessageID int                `bson:"message_id"`
	User      *tb.User           `bson:"user"`
	Timestamp int64              `bson:"timestamp"`
}

func InitMenu() {
	menu.Reply(
		menu.Row(btnRemind),
		menu.Row(btnList),
	)
}
