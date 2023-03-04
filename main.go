package main

import "echo-todo-server/src/router"

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
