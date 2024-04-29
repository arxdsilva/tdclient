package models

type (
	V4GetGuildResponse struct {
		Information APIInformation `json:"information"`
		Guild       Guild          `json:"guild"`
	}

	Guild struct {
		Name             string        `json:"name"`
		World            string        `json:"world"`
		LogoURL          string        `json:"logo_url"`
		Description      string        `json:"description"`
		Guildhalls       []GuildHall   `json:"guildhalls"`
		Active           bool          `json:"active"`
		Founded          string        `json:"founded"`
		OpenApplications bool          `json:"open_applications"`
		Homepage         string        `json:"homepage"`
		InWar            bool          `json:"in_war"`
		DisbandDate      string        `json:"disband_date"`
		DisbandCondition string        `json:"disband_condition"`
		PlayersOnline    int           `json:"players_online"`
		PlayersOffline   int           `json:"players_offline"`
		MembersTotal     int           `json:"members_total"`
		MembersInvited   int           `json:"members_invited"`
		Members          []GuildMember `json:"members"`
		Invites          interface{}   `json:"invites"`
	}

	GuildHall struct {
		Name      string `json:"name"`
		World     string `json:"world"`
		PaidUntil string `json:"paid_until"`
	}

	GuildMember struct {
		Name     string `json:"name"`
		Title    string `json:"title"`
		Rank     string `json:"rank"`
		Vocation string `json:"vocation"`
		Level    int    `json:"level"`
		Joined   string `json:"joined"`
		Status   string `json:"status"`
	}

	V4GetGuildsResponse struct {
		Guilds struct {
			Active    []GuildDescription `json:"active"`
			Formation []GuildDescription `json:"formation"`
			World     string             `json:"world"`
		} `json:"guilds"`
		Information APIInformation `json:"information"`
	}

	GuildDescription struct {
		Description string `json:"description"`
		LogoURL     string `json:"logo_url"`
		Name        string `json:"name"`
	}
)
