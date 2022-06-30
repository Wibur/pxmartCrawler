package main

import (
	"openCrawler/databases"
	"openCrawler/router"
)

func main() {
	databases.MySqlConnect()

	router.Register()
}
