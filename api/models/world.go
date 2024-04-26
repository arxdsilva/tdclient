package models

type V4GetWorldsResponse struct {
	Information struct {
		API       APIInfo   `json:"api"`
		Status    APIStatus `json:"status"`
		Timestamp string    `json:"timestamp"`
	} `json:"information"`
	Worlds TibiaInfo `json:"worlds"`
}

type V4GetWorldResponse struct {
	Information struct {
		API       APIInfo   `json:"api"`
		Status    APIStatus `json:"status"`
		Timestamp string    `json:"timestamp"`
	} `json:"information"`
	World World `json:"world"`
}

type APIInfo struct {
	Commit  string `json:"commit"`
	Release string `json:"release"`
	Version int    `json:"version"`
}

type APIStatus struct {
	Error    int    `json:"error"`
	HTTPCode int    `json:"http_code"`
	Message  string `json:"message"`
}

type TibiaInfo struct {
	PlayersOnline    int                     `json:"players_online"`
	RecordDate       string                  `json:"record_date"`
	RecordPlayers    int                     `json:"record_players"`
	RegularWorlds    []RegularWorldSimple    `json:"regular_worlds"`
	TournamentWorlds []TournamentWorldSimple `json:"tournament_worlds"`
}

type RegularWorldSimple struct {
	BattleyeDate        string `json:"battleye_date"`
	BattleyeProtected   bool   `json:"battleye_protected"`
	GameWorldType       string `json:"game_world_type"`
	Location            string `json:"location"`
	Name                string `json:"name"`
	PlayersOnline       int    `json:"players_online"`
	PremiumOnly         bool   `json:"premium_only"`
	PvpType             string `json:"pvp_type"`
	Status              string `json:"status"`
	TournamentWorldType string `json:"tournament_world_type"`
	TransferType        string `json:"transfer_type"`
}

type TournamentWorldSimple struct {
	BattleyeDate        string `json:"battleye_date"`
	BattleyeProtected   bool   `json:"battleye_protected"`
	GameWorldType       string `json:"game_world_type"`
	Location            string `json:"location"`
	Name                string `json:"name"`
	PlayersOnline       int    `json:"players_online"`
	PremiumOnly         bool   `json:"premium_only"`
	PvpType             string `json:"pvp_type"`
	Status              string `json:"status"`
	TournamentWorldType string `json:"tournament_world_type"`
	TransferType        string `json:"transfer_type"`
}

type World struct {
	BattleyeDate        string              `json:"battleye_date"`
	BattleyeProtected   bool                `json:"battleye_protected"`
	CreationDate        string              `json:"creation_date"`
	GameWorldType       string              `json:"game_world_type"`
	Location            string              `json:"location"`
	Name                string              `json:"name"`
	OnlinePlayers       []PlayerInformation `json:"online_players"`
	PlayersOnline       int                 `json:"players_online"`
	PremiumOnly         bool                `json:"premium_only"`
	PvpType             string              `json:"pvp_type"`
	RecordDate          string              `json:"record_date"`
	RecordPlayers       int                 `json:"record_players"`
	Status              string              `json:"status"`
	TournamentWorldType string              `json:"tournament_world_type"`
	TransferType        string              `json:"transfer_type"`
	WorldQuestTitles    []string            `json:"world_quest_titles"`
}

type PlayerInformation struct {
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Vocation string `json:"vocation"`
}
