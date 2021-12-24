package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	dbClient *mongo.Client
	dbCtx    context.Context
	dbCursor *mongo.Database
	dbCol    *mongo.Collection
	dbCancel context.CancelFunc
)

func messageFromStoreddReminder(stored StoredReminder) tb.Message {
	return tb.Message{ID: stored.MessageID, Chat: &tb.Chat{ID: stored.ChatID}}
}

func initDB(ctx context.Context) {
	dbCtx = ctx

	dbString := fmt.Sprintf("mongo://%s:%s", mongoHostname, mongoPort)
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(dbString))
	dbClient = client
	dbCursor = dbClient.Database("telegram")
	dbCol = dbCursor.Collection("reminders")

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
}

func getUserReminders(user *tb.User) ([]StoredReminder, error) {
	var reminders []StoredReminder
	cur, err := dbCol.Find(dbCtx, bson.M{"user.id": user.ID})
	if err != nil {
		return reminders, err
	}
	defer cur.Close(dbCtx)

	cur.All(dbCtx, &reminders)
	return reminders, nil
}

func loadStoredReminders() {
	var reminders []StoredReminder
	cur, err := dbCol.Find(dbCtx, bson.D{})
	if err != nil {
		log.Fatal("Error loading Saved Reminders")
	}
	defer cur.Close(dbCtx)
	cur.All(dbCtx, &reminders)

	for i := range reminders {
		timestamp := time.Unix(reminders[i].Timestamp, 0)
		duration := timestamp.Sub(time.Now())
		fmt.Println(duration)

	}
}

func getStoredReminders(id primitive.ObjectID) (StoredReminder, error) {
	reminder := StoredReminder{}

	return reminder, nil
}
