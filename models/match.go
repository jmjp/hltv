package models

import "time"

type UpcomingMatche struct {
	Id     int
	TeamA  *Team
	TeamB  *Team
	Format string
	Event  *string
	Live   bool
	Date   time.Time
}
