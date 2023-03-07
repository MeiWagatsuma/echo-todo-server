package model

import (
	"echo-todo-server/src/env"
	"fmt"
	"log"
	"time"
)

func CreateSessionTable() (err error) {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS sessions (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		session_key VARCHAR(%d) UNIQUE NOT NULL,
		expires_at TIMESTAMPTZ NOT NULL
	)`, env.SESSION_KEY_LENGTH)
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

type Session struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	SessionKey string    `json:"session_key"`
	ExpiresAt  time.Time `json:"expires_at"`
}
