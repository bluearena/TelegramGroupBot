package main

import "time"

func getUser(userId int) User {
	var u User
	db.Find(&u, "id = ?", userId)
	return u
}

func insertUser(userId int) {
	u := User{userId, 0, time.Time{}, false}
	db.Create(&u)
}

func getViolationCount(userId int) int {
	u := getUser(userId)
	if u == (User{}) {
		insertUser(userId)
		return 0
	}
	return u.Violations
}

func incrementViolationCount(userId int) {
	var u User
	db.Find(&u, "id = ?", userId)
	u.Violations++
	db.Save(u)
}

func setIsBanned(userId int) {
	var u User
	db.Find(&u, "id = ?", userId)
	u.IsBanned = true
	db.Save(u)
}

func getTimeToCheckName(userId int) time.Time {
	u := getUser(userId)
	if u == (User{}) {
		insertUser(userId)
		return time.Time{}
	}
	return u.TimeToCheckName
}

func setTimeToCheckName(userId int) {
	var u User
	db.Find(&u, "id = ?", userId)
	u.TimeToCheckName = time.Now().Add(72 * time.Hour)
	db.Save(u)
}
