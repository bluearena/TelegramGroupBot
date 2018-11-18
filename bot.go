package main

import (
	"log"
	"os"
	"io"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	bot                *tgbotapi.BotAPI
	config             Configuration
	db                 *gorm.DB
	phrases            Phrases
)

func main() {
	initLog()
	initConfig()
	initPhrases()
	initDB()

	var err error
	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Print("ERROR: ")
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil{
			handleMessage(update.Message)
		}
	}

}

func initLog() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Print("ERROR: ")
		log.Panic(err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func initConfig() {
	readJson(&config, "config.json")
}

func initPhrases() {
	readJson(&phrases, "strings.json")
}

func initDB() {
	var err error
	db, err = gorm.Open("sqlite3", "data.db")
	//db, err = gorm.Open("mysql", config.DbArgs)
	if err != nil {
		log.Print("********** ERROR: " + err.Error())
		log.Panic("Failed to connect database")
	} else {
		log.Print("Opened DB")
	}
	db.LogMode(true)
	log.Print("Set LogMode")
	db.AutoMigrate(&User{})
	log.Print("Migrated")
}
