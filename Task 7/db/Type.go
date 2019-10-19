package db

import (
	"pq-master"
	"time"
)

//This struct is for "client" (Main.go)
type Task struct {
	Id           int
	Text         string
	CreateTime   time.Time
	CompleteTime pq.NullTime
}
