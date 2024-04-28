package models

type (
	V4GetBoostableBossesResponse struct {
		BoostableBosses BoostableBosses `json:"boostable_bosses"`
		Information     APIInformation  `json:"information"`
	}

	BoostableBosses struct {
		BoostableBossList []BoostableBoss `json:"boostable_boss_list"`
		Boosted           Boosted         `json:"boosted"`
	}

	BoostableBoss struct {
		Featured bool   `json:"featured"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
	}

	Boosted struct {
		Featured bool   `json:"featured"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
	}
)
