package main

import (
	"fmt"
	"os"
	"github.com/8tomat8/SSU-Golang-252-Chat/loger"

	"github.com/YAWAL/GoTraining/database/redis"
)

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}


func main() {
	conn, err := redis.GetRedisConn()
	if err != nil {
		loger.Log.Error(err)
	}

	resp := conn.Cmd("ping")
	fmt.Printf("resi value is %v , type is %T \n", resp, resp )
	// select database
	r := conn.Cmd("select", 0)
	errHndlr(r.Err)

	r = conn.Cmd("echo", "Hello world!")
	errHndlr(r.Err)

	r = conn.Cmd("set", "mykey0", "myval0")
	errHndlr(r.Err)

	s, err := conn.Cmd("get", "mykey0").Str()
	errHndlr(err)
	fmt.Println("mykey0:", s)

	myhash := map[string]string{
		"mykey1": "myval1",
		"mykey2": "myval2",
		"mykey3": "myval3",
	}

	// Alternatively:
	// c.Cmd("mset", "mykey1", "myval1", "mykey2", "myval2", "mykey3", "myval3")
	r = conn.Cmd("mset", myhash)
	errHndlr(r.Err)


	r = conn.Cmd("SET", "arg1", 5)
	errHndlr(r.Err)
	s, err = conn.Cmd("GET", "arg1").Str()
	//errHndlr(r.Err)
	fmt.Println("arg1:", s)







	//r = conn.Cmd("zadd", "rertytrh")
	//r = conn.Cmd("zadd", "yjgjjg")


	errHndlr(r.Err)



}
