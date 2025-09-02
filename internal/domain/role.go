package domain

import "time"

type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleSimplified struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
