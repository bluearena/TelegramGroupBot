package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"time"
)

func handleViolation(message *tgbotapi.Message) {
	deleteMessage(config.GroupId, message.MessageID)
	count := getViolationCount(message.From.ID)
	if count < 3 {
		makeWarning(message)
	} else {
		banUser(message)
	}
}

func makeWarning(message *tgbotapi.Message) {
	incrementViolationCount(message.From.ID)
	sendMessage(config.GroupId, getWarningWordsMessage(message.From.FirstName + " " + message.From.LastName), nil)
}

func banUser(message *tgbotapi.Message) {
	kickUser(config.GroupId, message.From.ID)
	setIsBanned(message.From.ID)
	sendMessage(config.GroupId, getUserBannedMessage(message.From.FirstName + " " + message.From.LastName), nil)
}

func handleWrongName(message *tgbotapi.Message) {
	t := getTimeToCheckName(message.From.ID)
	if t == (time.Time{}) {
		setTimeToCheckName(message.From.ID)
		sendMessage(config.GroupId, getWarningNameMessage(message.From.FirstName + " " + message.From.LastName), nil)
	} else {
		if time.Now().After(t){
			kickUser(config.GroupId, message.From.ID)
			setIsBanned(message.From.ID)
			sendMessage(config.GroupId, getUserBannedMessage(message.From.FirstName + " " + message.From.LastName), nil)
		}
	}
}
