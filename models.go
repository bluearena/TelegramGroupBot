package main

type Configuration struct {
	BotToken    string
	DbArgs      string
	AdminIds    []int64
	GroupId     int64
	BannedWords []string
}

type UserViolations struct {
	Id         int `gorm:"primary_key"`
	Violations int
	IsBanned   bool
}
