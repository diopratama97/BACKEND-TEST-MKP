package main

import (
	"backend-test-mkp/db"
	route "backend-test-mkp/routes"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":80"))
}
