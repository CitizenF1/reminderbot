package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	tb "gopkg.in/tucnak/telebot.v2"
)

type StoredReminder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ChatID    int64              `bson:"chat_id"`
	MessageID int                `bson:"message_id"`
	User      *tb.User           `bson:"user"`
	Timestamp int64              `bson:"timestamp"`
}
