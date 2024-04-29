package models

type (
	V4GetKillStatisticsResponse struct {
		Information    APIInformation `json:"information"`
		Killstatistics struct {
			Entries []KillEntry      `json:"entries"`
			Total   TotalInformation `json:"total"`
			World   string           `json:"world"`
		} `json:"killstatistics"`
	}

	KillEntry struct {
		LastDayKilled         int    `json:"last_day_killed"`
		LastDayPlayersKilled  int    `json:"last_day_players_killed"`
		LastWeekKilled        int    `json:"last_week_killed"`
		LastWeekPlayersKilled int    `json:"last_week_players_killed"`
		Race                  string `json:"race"`
	}

	TotalInformation struct {
		LastDayKilled         int `json:"last_day_killed"`
		LastDayPlayersKilled  int `json:"last_day_players_killed"`
		LastWeekKilled        int `json:"last_week_killed"`
		LastWeekPlayersKilled int `json:"last_week_players_killed"`
	}
)
