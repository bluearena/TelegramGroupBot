package main

func getUserViolations(userId int) *UserViolations {
	var uv *UserViolations
	db.Find(uv, "id = ?", userId)
	return uv
}

func insertUser(userId int){
	uv := UserViolations{userId, 0, false}
	db.Create(&uv)
}

func getViolationCount(userId int) int{
	uv := getUserViolations(userId)
	if uv == nil{
		insertUser(userId)
		return 0
	}
	return uv.Violations
}

func incrementViolationCount(userId int){
	var uv *UserViolations
	db.Find(uv, "id = ?", userId)
	uv.Violations++
	db.Save(uv)
}

func setIsBanned(userId int){
	var uv *UserViolations
	db.Find(uv, "id = ?", userId)
	uv.IsBanned = true
	db.Save(uv)
}