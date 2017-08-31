package main

import (
	"fmt"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"
	"github.com/YAWAL/GoTraining/database/redis"
)

func main() {

	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Errorf("error 1")
	}
	resp := conn.Cmd("Set", "authos", "zaz", "bmw", "toyota")
	resp.Str()
	s, err := conn.Cmd("SADD", "authos", "zaz", "bmw", "toyota").Str()
	if err != nil {
		loger.Log.Errorf("error 2", err)
	}
	fmt.Printf(s)
}
