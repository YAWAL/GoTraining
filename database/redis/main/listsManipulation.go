package main

import (
	"fmt"
	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database/redis"
)

func main()  {

	conn, err := redis.GetRedisConn()
	if err != nil{
		//err handling
	}

	ls, err := conn.Cmd("mget", "mykey1", "mykey2", "mykey3").List()

	if err != nil{
		//err handling
	}
	fmt.Println("mykeys values:", ls)

	fmt.Println("====================================================")
	conn.Cmd("set", "1","one", "2", "two").List()
	if err != nil{
		//err handling
	}

	conn.Cmd("rpush", "two","four").List()

	ls2, err := conn.Cmd("mget", "0", "-1" ).List()
	if err != nil{
		fmt.Println("erra")
	}
	fmt.Println("mykeys values:", ls2)

}
