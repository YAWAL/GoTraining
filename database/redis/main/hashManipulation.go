package main

import (
	"github.com/YAWAL/GoTraining/database/redis"
)

func main() {

	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Error(err)
	}



}
