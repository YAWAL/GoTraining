package main

import (
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database/redis"
	"fmt"
)

func main() {

	conn, err := redis.GetRedisConn()
	if err != nil {

	}

	s, err := conn.Cmd("SADD", "authos", "zaz", "bmw", "toyota").Str()
	if err != nil {

	}
	fmt.Println(s)
}
