package main

import (
	"strconv"
	"strings"
	"time"

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
		botInstance.Send(m.Chat, "Не правильная единица времени!")
		return
	}

	if m.ReplyTo == nil {
		botInstance.Send(m.Chat, "Нет сообщения для пересылки !")
		return
	}

	go confirmReminderSet(wait, m.Chat)
	go forwardMessageAfterDelay(wait, m.Sender, m.ReplyTo)
}

//TODO
//отмена напоминаний
func cancelHandler(m *tb.Message) {

}

func listHandler(m *tb.Message) {
	remiders, done := reminderHelper(m)
	if !done {
		textArray := []string{"У вас " + strconv.Itoa(len(remiders)) + " Активных напоминаний:"}
		for i := range remiders {
			textArray = append(textArray, time.Unix(remiders[i].Timestamp, 0).String())
		}
		text := strings.Join(textArray, "\n")
		botInstance.Send(m.Sender, text)
	}
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
			botInstance.Send(m.Sender, "Нет совпадений!...")
		} else {
			go confirmReminderSet(wait, m.Sender)
			go forwardMessageAfterDelay(wait, m.Sender, waitingMessage)
		}

		delete(currentLimboUsers, m.Sender.ID)
	} else {
		currentLimboUsers[m.Sender.ID] = m
		botInstance.Send(m.Sender, "Когда мне напомнить?")
	}
}
