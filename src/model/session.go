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
		token VARCHAR(%d) PRIMARY KEY NOT NULL,
		user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
		expires_at TIMESTAMPTZ NOT NULL
	)`, env.SESSION_KEY_LENGTH*6)
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

type Session struct {
	Token     string    `json:"token"`
	UserId    string    `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (s *Session) Exists(userId string) (exists bool, err error) {
	query := "SELECT EXISTS(SELECT 1 FROM sessions WHERE user_id=$1)"

	if err = Db.QueryRow(query, userId).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (s *Session) Generate() (err error) {
	token, err := lib.GenerateRandomBase64String(env.SESSION_KEY_LENGTH)
	if err != nil {
		log.Println("GenerateRandomBase64String failed: ", err)
		return err
	}

	s.ExpiresAt = time.Now().Add(time.Second * 30)

	query := "INSERT INTO sessions(token, user_id, expires_at) VALUES($1, $2, $3)"
	if _, err = Db.Exec(query, token, s.UserId, s.ExpiresAt); err != nil {
		log.Println("Insert session key failed: ", err)
		return err
	}

	// If an error occurs when inserting a session,
	// I don't want to put a token in the instance,
	// so I'm assuming token here
	s.Token = token
	return nil
}

func (s Session) DeleteExpiredSessions() error {
	query := `DELETE FROM sessions WHERE expires_at < NOW()`
	_, err := Db.Exec(query)
	if err != nil {
		return err
	}
	log.Println("Expired sessions were deleted!")
	return nil
}
