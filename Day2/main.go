package main

import (
	"Day2/config"
	"Day2/routes"
)

func main() {
	config.InitDb()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
