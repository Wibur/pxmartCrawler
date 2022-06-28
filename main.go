package main

import (
	"log"
	"openCrawler/databases"
	"openCrawler/router"
)

func main() {
	go func() {
		log.Println("db connect")
		databases.MySqlConnect()
	}()
	router.Register()
}
