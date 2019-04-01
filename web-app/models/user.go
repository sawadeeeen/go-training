package models

import (
	"regexp"
	"time"
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAT time.Time
	DeletedAt time.Time
}
