package model

import "log"

func CreateUserTable() (err error) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`
	_, err = Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
