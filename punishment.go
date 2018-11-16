package main

import "github.com/go-telegram-bot-api/telegram-bot-api"

func handleViolation(message *tgbotapi.Message){
	deleteMessage(config.GroupId, message.MessageID)
	count := getViolationCount(message.From.ID)
	if count < 3{
		makeWarning(message)
	}else {
		banUser(message)
	}
}

func makeWarning(message *tgbotapi.Message){
	incrementViolationCount(message.From.ID)
	sendMessage(config.GroupId, "Warning", nil)
}

func banUser(message *tgbotapi.Message){
	kickUser(config.GroupId, message.From.ID)
	setIsBanned(message.From.ID)
	sendMessage(config.GroupId, "User was banned", nil)
}