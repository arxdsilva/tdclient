package models

type (
	V4GetHighScores struct {
		Highscores struct {
			Category      string          `json:"category"`
			HighscoreAge  int             `json:"highscore_age"`
			HighscoreList []Highscore     `json:"highscore_list"`
			HighscorePage PageInformation `json:"highscore_page"`
			Vocation      string          `json:"vocation"`
			World         string          `json:"world"`
		} `json:"highscores"`
		Information APIInformation `json:"information"`
	}

	Highscore struct {
		Level    int    `json:"level"`
		Name     string `json:"name"`
		Rank     int    `json:"rank"`
		Title    string `json:"title"`
		Value    int    `json:"value"`
		Vocation string `json:"vocation"`
		World    string `json:"world"`
	}

	PageInformation struct {
		CurrentPage  int `json:"current_page"`
		TotalPages   int `json:"total_pages"`
		TotalRecords int `json:"total_records"`
	}
)
