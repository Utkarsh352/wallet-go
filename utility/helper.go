package utility

import "time"

func FormatTimestamp(t time.Time) string {
	return t.Format("15:04:05 02/01/2006")
}
