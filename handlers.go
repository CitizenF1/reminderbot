package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func setHandlers() {
	botInstance.Handle("/start", startHandler)
	botInstance.Handle(btnRemind, remindHandler)
	botInstance.Handle(btnList, listHandler)

	botInstance.Handle(tb.OnText, onTextHandler)
}

func startHandler(m *tb.Message) {
	botInstance.Send(m.Sender, "Привет", menu)
}

func remindHandler(m *tb.Message) {
	wait, err := getWaitTime(m.Payload)
	if m.Payload == "" || err != nil {
		botInstance.Send(m.Chat, "No valid time units found!")
		return
	}

	if m.ReplyTo == nil {
		botInstance.Send(m.Chat, "No message to farward!")
		return
	}

	go confirmReminderSet(wait, m.Chat)
	go forwardMessageAfterDelay(wait, m.Sender, m.ReplyTo)
}

func listHandler(m *tb.Message) {

}

func onTextHandler(m *tb.Message) {
	// if privateMessageHelper(m, true) {
	// 	return
	// }

	// log.Println("Private Message!")

	if waitingMessage, ok := currentLimboUsers[m.Sender.ID]; ok {
		// Already waiting
		wait, err := getWaitTime(m.Text)
		if err != nil {
			botInstance.Send(m.Sender, "No valid match! Aborting...")
		} else {
			go confirmReminderSet(wait, m.Sender)
			go forwardMessageAfterDelay(wait, m.Sender, waitingMessage)
		}

		delete(currentLimboUsers, m.Sender.ID)
	} else {
		currentLimboUsers[m.Sender.ID] = m
		botInstance.Send(m.Sender, "When should I remind you?")
	}
}
