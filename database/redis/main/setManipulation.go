package main

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database/redis"
	"fmt"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
)

func main() {

	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Errorf("error")
	}

	s, err := conn.Cmd("SADD", "authos", "zaz", "bmw", "toyota").Str()
	if err != nil {
		loger.Log.Errorf("error")
	}
	fmt.Println(s)
}
