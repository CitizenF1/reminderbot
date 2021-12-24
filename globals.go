package main

import tb "gopkg.in/tucnak/telebot.v2"

var (
	envBotToken   string
	envListenPort string
	// envRootPublicURL      string
	botMode               string
	publicURL             string
	rootTelegramMethodURL string
	mongoHostname         string
	mongoPort             string
	currentLimboUsers     map[int]*tb.Message = make(map[int]*tb.Message)
)
