package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Redis server 6379 portunda docker server

func main() {
	fmt.Println("Hello world!")
	conn, err := redis.Dial("tcp", "localhost:6379")
	checkError(err)
	defer conn.Close()
	// SET
	_, err = conn.Do("SET", "something", "something else")
	// GET
	reply, err := redis.String(conn.Do("GET", "something"))

	fmt.Println(reply)

	// conn.Do() direkt cli'daki gibi komut yazmaya yariyormus DELETE falan da calisiyor
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
