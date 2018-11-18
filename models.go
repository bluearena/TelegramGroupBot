package main

import "time"

type Configuration struct {
	BotToken    string
	DbArgs      string
	AdminIds    []int64
	GroupId     int64
	BannedWords []string
	BannedNames []string
}

type User struct {
	Id         int `gorm:"primary_key"`
	Violations int
	TimeToCheckName time.Time
	IsBanned   bool
}

type Phrases struct {
	Welcomes []string
	WarningName, WarningWords, UserBanned string
}