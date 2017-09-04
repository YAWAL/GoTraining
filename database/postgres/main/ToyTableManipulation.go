package main

import "github.com/YAWAL/GoTraining/database/postgres"
import "github.com/8tomat8/SSU-Golang-252-Chat/loger"

type Toy struct {
	id      int `gorm:"primary_key"`
	name    string
	isExist bool
}

func main() {
	db, err := postgres.GetPostgresConnection()

	if err != nil {
		loger.Log.Errorf("error: ", err)
	}

	user := User{"vasyl", 45}
	db.NewRecord(user) // => returns `true` as primary key is blank

	db.Create(&user)


}
