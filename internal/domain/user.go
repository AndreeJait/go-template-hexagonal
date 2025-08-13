package domain

import "time"

type User struct {
	ID       int64
	Email    string
	Password string
	CreateAt time.Time
	UpdateAt time.Time
}
