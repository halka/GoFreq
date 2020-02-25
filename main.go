package main

import (
	"./db"
	"./router"
)

func main() {
	db.Init()
	defer db.Close()
	router.Setup()
}
