package db

import (
	"github.com/lib/pq"
	"time"
)

type Task struct { 
	Id           int
	Text         string
	CreateTime   time.Time
	CompleteTime pq.NullTime
}
