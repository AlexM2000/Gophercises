package db

import (
	"time"
)

type Task struct {
	ID         int
	Text       string
	CreateTime time.Time
	Complete   bool
}
