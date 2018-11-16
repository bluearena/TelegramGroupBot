package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func handleMessage(message *tgbotapi.Message){
	if message.NewChatMembers != nil{
		handleNewMembers(message.NewChatMembers)
	}else if message.Text != ""{
		handleText(message)
	}
}

func handleNewMembers(newMembers *[]tgbotapi.User){
	for _, member := range *newMembers{
		sendMessage(config.GroupId, "Welcome, " + member.FirstName, nil)
	}
}

func handleText(message *tgbotapi.Message){
	if message.Entities != nil{
		if isDeleted := handleEntities(message); isDeleted {
			return
		}
	}
	handleWords(message)
}

func handleEntities(message *tgbotapi.Message) bool {
	for _, entity := range *message.Entities{
		if (entity.Type == "url" || entity.Type == "text_link") && !containsInt64(config.AdminIds, int64(message.From.ID)){
			handleViolation(message)
			return true
		}
	}
	return false
}

func handleWords(message *tgbotapi.Message){
	for _, word := range config.BannedWords{
		if strings.Contains(message.Text, word){
			handleViolation(message)
		}
	}
}