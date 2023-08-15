package models

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TeamInfo struct {
	Team
	LogoURL          string         `json:"logo"`
	Country          string         `json:"country"`
	WorldRanking     uint8          `json:"world_ranking"`
	WeeksInTop30     int            `json:"weeks_in_top30"`
	AveragePlayerAge float32        `json:"average_player_age"`
	CoachName        string         `json:"coach_name"`
	Lineup           *[]TeamPlayers `json:"lineup"`
}

type TeamPlayers struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Status     string  `json:"status"`
	TimeOnTeam float32 `json:"time_on_team"`
	MapsPlayed int     `json:"maps_Played"`
	Rating     float32 `json:"rating"`
}
