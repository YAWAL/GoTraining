package main

import (
	"fmt"
	"github.com/YAWAL/GoTraining/database/redis"
)

func main()  {

	conn, err := redis.GetRedisConn()
	if err != nil{

	}

	mylist := []string{"one", "two", "three"}

	// Alternativaly:
	// c.Cmd("rpush", "mylist", "foo", "bar", "qux")
	r := conn.Cmd("rpush", "mylist", mylist)
	if r.Err != nil {
	}

	mylist, err = conn.Cmd("lrange", "mylist", 0, -1).List()
	if r.Err != nil {
	}

	fmt.Println("mylist:", mylist)
}
