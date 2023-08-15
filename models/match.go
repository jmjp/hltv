package models

import "time"

type UpcomingMatche struct {
	Id     int       `json:"id"`
	TeamA  *Team     `json:"team_a"`
	TeamB  *Team     `json:"team_b"`
	Format string    `json:"format"`
	Event  *string   `json:"event"`
	Live   bool      `json:"live"`
	Date   time.Time `json:"date"`
}
