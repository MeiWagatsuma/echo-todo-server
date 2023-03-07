package model

import (
	"echo-todo-server/src/env"
	"echo-todo-server/src/lib"
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

func (s *Session) generate(userId string) (sessionKey string, err error) {
	query := "INSERT INTO sessions(user_id, session_key, expires_at) VALUES($1, $2, $3)"

	if sessionKey, err = lib.GenerateRandomBase64String(env.SESSION_KEY_LENGTH); err != nil {
		log.Println("GenerateRandomBase64String failed: ", err)
		return "", err
	}

	expiresAt := time.Now().Add(time.Second * 30)

	if _, err = Db.Exec(query, userId, sessionKey, expiresAt); err != nil {
		log.Println("Insert session key failed")
		return "", err
	}

	return sessionKey, err
}
