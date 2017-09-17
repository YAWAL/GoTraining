package main

import (
	"fmt"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"

	"github.com/YAWAL/GoTraining/database/redis"
)

func main() {
	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Error(err)
	}

	conn.Cmd("INCR", "arg1")
	s, err := conn.Cmd("GET", "arg1").Str()
	fmt.Println("arg1:", s)

	s, err =  conn.Cmd("set", "stringKey","justAstring").Str()
	fmt.Println("stringKey:", s)

	conn.Cmd("APPEND", "stringKey","appended")

	s, err = conn.Cmd("GET", "stringKey").Str()
	fmt.Println("stringKey:", s)
}
