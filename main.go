package main

import (
	"openCrawler/router"
	"openCrawler/databases"
)

func main() {
	router.Register()

	databases.MySqlConnect()
}
