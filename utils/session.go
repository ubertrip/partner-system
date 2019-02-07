package utils

import (
	"time"
)

func Midnight() time.Time {
	tm := time.Now()
	year, month, day := tm.Date()
	return time.Date(year, month, day+1, 0, 0, 0, 0, tm.Location())
}
