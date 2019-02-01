package format

import "time"

// Time formats time as YYYY-MM-DD HH:MM:SS
func Time(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
