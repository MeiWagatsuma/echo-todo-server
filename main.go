package main

import (
	"echo-todo-server/src/model"
	"echo-todo-server/src/router"
	"fmt"
)

func main() {
	e := router.New()

	fmt.Println(model.Db)

	e.Logger.Fatal(e.Start(":8080"))
}
