package main

import "github.com/YAWAL/GoTraining/database/postgres"
import "github.com/8tomat8/SSU-Golang-252-Chat/loger"

type User struct {
	Name string `gorm:"primary_key"`
	Age int
}

type Toy struct {
	Id      int `gorm:"primary_key"`
	Name    string
	IsExist bool
}

func main() {
	db, err := postgres.GetPostgresConnection()

	if err != nil {
		loger.Log.Errorf("error: ", err)
	}

	// Create table for model `User` and 'Toy'
	db.CreateTable(&User{})
	db.CreateTable(&Toy{})

}
