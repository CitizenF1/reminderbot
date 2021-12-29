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

	dbString := fmt.Sprintf("mongodb://%s:%s", mongoHostname, mongoPort)
	fmt.Println(dbString, "dbString")
	clientOption := options.Client().ApplyURI(dbString)
	client, err := mongo.Connect(dbCtx, clientOption)
	if err != nil {
		log.Fatal(err, "++++++++=")
	}
	dbClient = client
	dbCursor = dbClient.Database("telegram")
	fmt.Println(dbCursor, "DBCURSOR")
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

func getStoredRemindersID(id primitive.ObjectID) (StoredReminder, error) {
	reminder := StoredReminder{}
	res := dbCol.FindOne(dbCtx, bson.M{"_id": id})
	err := res.Decode(&reminder)
	if err != nil {
		log.Println("Unable to load DB message from ID")
		return reminder, err
	}
	return reminder, nil
}

func storeMessageIntoDB(m *tb.Message, recipient *tb.User, timestamp int64) primitive.ObjectID {
	// Return ObjectId
	storedReminder := StoredReminder{ChatID: m.Chat.ID, MessageID: m.ID, User: recipient, Timestamp: timestamp}
	res, err := dbCol.InsertOne(dbCtx, storedReminder)
	if err != nil {
		log.Panic(err)
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id
}

func removeMessageFromDB(id primitive.ObjectID) int64 {
	res, err := dbCol.DeleteMany(dbCtx, bson.M{})
	if err != nil {
		log.Panic(err)
	}
	return res.DeletedCount
}
