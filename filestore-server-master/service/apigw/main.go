package main

import (
	"filestore-server/service/apigw/route"

	"log"
	"runtime/debug"
)

func main() {
	log.Default().Println("agw服务启动")
	defer func() {
		log.Default().Println("agw服务异常")
		if r := recover(); r != nil {
			log.Default().Println("Recovered from panic:", r)
			log.Default().Printf("Stack trace:\n%s\n", debug.Stack())
		}
	}()
	r := route.Router()
	r.Run(":8080")
}
