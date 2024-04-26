package models

type (
	V4GetCharacterResponse struct {
		Character   Character      `json:"character"`
		Information APIInformation `json:"information"`
	}

	Character struct {
		AccountBadges      []Badge            `json:"account_badges"`
		AccountInformation AccountInformation `json:"account_information"`
		Achievements       []Achievement      `json:"achievements"`
		Character          CharacterInfo      `json:"character"`
		Deaths             []Death            `json:"deaths"`
		DeathsTruncated    bool               `json:"deaths_truncated"`
		OtherCharacters    []OtherCharacters  `json:"other_characters"`
	}

	Badge struct {
		Description string `json:"description"`
		IconURL     string `json:"icon_url"`
		Name        string `json:"name"`
	}

	AccountInformation struct {
		Created      string `json:"created"`
		LoyaltyTitle string `json:"loyalty_title"`
		Position     string `json:"position"`
	}

	Achievement struct {
		Grade  int    `json:"grade"`
		Name   string `json:"name"`
		Secret bool   `json:"secret"`
	}

	CharacterInfo struct {
		AccountStatus     string      `json:"account_status"`
		AchievementPoints int         `json:"achievement_points"`
		Comment           string      `json:"comment"`
		DeletionDate      string      `json:"deletion_date"`
		FormerNames       []string    `json:"former_names"`
		FormerWorlds      []string    `json:"former_worlds"`
		Guild             Guild       `json:"guild"`
		Houses            []HouseInfo `json:"houses"`
		LastLogin         string      `json:"last_login"`
		Level             int         `json:"level"`
		MarriedTo         string      `json:"married_to"`
		Name              string      `json:"name"`
		Position          string      `json:"position"`
		Residence         string      `json:"residence"`
		Sex               string      `json:"sex"`
		Title             string      `json:"title"`
		Traded            bool        `json:"traded"`
		UnlockedTitles    int         `json:"unlocked_titles"`
		Vocation          string      `json:"vocation"`
		World             string      `json:"world"`
	}

	Guild struct {
		Name string `json:"name"`
		Rank string `json:"rank"`
	}

	HouseInfo struct {
		Houseid int    `json:"houseid"`
		Name    string `json:"name"`
		Paid    string `json:"paid"`
		Town    string `json:"town"`
	}

	Death struct {
		Assists []DeathAssist `json:"assists"`
		Killers []DeathKiller `json:"killers"`
		Level   int           `json:"level"`
		Reason  string        `json:"reason"`
		Time    string        `json:"time"`
	}

	DeathAssist struct {
		Name   string `json:"name"`
		Player bool   `json:"player"`
		Summon string `json:"summon"`
		Traded bool   `json:"traded"`
	}

	DeathKiller struct {
		Name   string `json:"name"`
		Player bool   `json:"player"`
		Summon string `json:"summon"`
		Traded bool   `json:"traded"`
	}

	OtherCharacters struct {
		Deleted  bool   `json:"deleted"`
		Main     bool   `json:"main"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Status   string `json:"status"`
		Traded   bool   `json:"traded"`
		World    string `json:"world"`
	}
)
