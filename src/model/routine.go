package model

import (
	"echo-todo-server/src/env"
	"fmt"
	"log"
	"time"
)

var session Session

// GoRoutine is for periodic execution
func GoRoutine() {
	fmt.Println("GoRoutine started!")
	for {
		if err := session.DeleteExpiredSessions(); err != nil {
			log.Println("deleteExpiredSessions was failed: ", err)
		}

		time.Sleep(time.Hour * time.Duration(env.SESSION_EXPIRATION_CHECK_INTERVAL_HOURS))
	}
}
