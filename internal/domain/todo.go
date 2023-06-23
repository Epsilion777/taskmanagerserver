package domain

import "time"

type Todo struct {
	ID       int
	UserID   string
	Title    string
	Progress int
	DateEnd  time.Time
}
