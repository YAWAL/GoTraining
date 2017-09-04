package main

import (
	"github.com/YAWAL/GoTraining/database/postgres"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

func main()  {
	db, err := postgres.GetPostgresConnection()

	if err != nil{
		loger.Log.Errorf("error: ", err)
	}

	ok := db.HasTable("toys")

	loger.Log.Info("table exist: ", ok)
}
