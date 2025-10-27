package domain

import (
	"time"
)

type Log struct {
	Timestamp time.Time
	Level string
	Message string
	Service string
}

// info
// error
// debug
// trace/verbose
// warn
// fatal
