package models

type PlayerSummary struct {
	Player
	Image   *string        `json:"image"`
	Team    *Team          `json:"team"`
	Stats   []StatEntry    `json:"stats"`
	Metrics *PlayerMetrics `json:"metrics"`
}

type StatEntry struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string
}

type PlayerMetrics struct {
	TotalKills      int     `json:"total_kills"`
	HeadshotPercent float32 `json:"headshot_percent"`
	TotalDeaths     int     `json:"total_deaths"`
	KDRatio         float32 `json:"kd_ratio"`
	DamagePerRound  float32 `json:"damage_per_round"`
	GrenadeDmgRound float32 `json:"grenade_dmg_round"`
	MapsPlayed      int     `json:"maps_played"`
	RoundsPlayed    int     `json:"rounds_played"`
	KillsPerRound   float32 `json:"kills_per_round"`
	AssistsPerRound float32 `json:"assists_per_round"`
	DeathsPerRound  float32 `json:"deaths_per_round"`
	SavedByTeammate float32 `json:"saved_by_team_mate"`
	SavedTeammates  float32 `json:"saved_team_mates"`
	Rating1_0       float32 `json:"ratio1.0"`
}
