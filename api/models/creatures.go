package models

type (
	V4GetCreaturesResponse struct {
		Creatures struct {
			Boosted      BoostedInfo    `json:"boosted"`
			CreatureList []CreatureInfo `json:"creature_list"`
		} `json:"creatures"`
		Information APIInformation `json:"information"`
	}

	CreatureInfo struct {
		Featured bool   `json:"featured"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
		Race     string `json:"race"`
	}

	BoostedInfo struct {
		Featured bool   `json:"featured"`
		ImageURL string `json:"image_url"`
		Name     string `json:"name"`
		Race     string `json:"race"`
	}
)
