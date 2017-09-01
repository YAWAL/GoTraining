package main

import (
	"github.com/YAWAL/GoTraining/database/redis"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"

)

func main() {

	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Error(err)
	}





}
