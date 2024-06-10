package models

import "time"

type Todo struct {
	ID           string    `db:"id"`
	Topic        string    `db:"topic"`
	Description  string    `db:"description"`
	CreationTime time.Time `db:"create_time"`
}
