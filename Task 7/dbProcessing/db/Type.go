package db

import (
	"pq-master"
	"time"
)

type Task struct { //This struct is for "server" (dbProcessing.go)
	Id           int
	Text         string
	CreateTime   time.Time
	CompleteTime pq.NullTime
}
