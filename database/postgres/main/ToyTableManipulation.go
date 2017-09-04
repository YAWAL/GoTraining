package main

import "github.com/YAWAL/GoTraining/database/postgres"
import "github.com/8tomat8/SSU-Golang-252-Chat/loger"

type Toy struct {
	id      int `gorm:"primary_key"`
	name    string
	isExist bool
}


type User struct {
	Name string `gorm:"primary_key"`
	Age int
}

func main() {
	db, err := postgres.GetPostgresConnection()

	if err != nil {
		loger.Log.Errorf("error: ", err)
	}

	user :=User{"eyrye", 22}
	db.NewRecord(user) // => returns `true` as primary key is blank
	db.Create(&user)

	toy := Toy{"b", true}
	db.NewRecord(toy)
	db.Create(&toy)

}
