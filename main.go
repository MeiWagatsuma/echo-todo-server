package main

import "echo-todo-server/router"

func main() {
	e := router.New()

	e.Logger.Fatal(e.Start(":8080"))
}
