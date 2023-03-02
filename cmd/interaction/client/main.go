package main

import (
	"log"
	"runedance_douyin/pkg/tools"
)

func main() {
	token, _ := tools.GenToken("123", 2)
	log.Println(token)
}
