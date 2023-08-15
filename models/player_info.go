package models

type PlayerInfo struct {
	Player
	Image       *string          `json:"image"`
	CurrentTeam *Team            `json:"team"`
	Top20       *string          `json:"presence_in_top20"` // Usando um ponteiro para permitir valor nulo
	MajorWins   *int16           `json:"major_wins"`
	Stats       *PlayerInfoStats `json:"stats"`
}

type PlayerInfoStats struct {
	Rating         float32 `json:"rating"`
	KillsPerRound  float32 `json:"kills_per_round"`
	Headshots      float32 `json:"headshots"`
	MapsPlayed     int     `json:"maps_played"`
	DeathsPerRound float32 `json:"deaths_per_round"`
	RoundsContrib  float32 `json:"rounds_contributed"`
}
