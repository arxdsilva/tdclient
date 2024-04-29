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

	V4GetCreatureByNameResponse struct {
		Creature    CreatureComplete `json:"creature"`
		Information APIInformation   `json:"information"`
	}

	CreatureComplete struct {
		BeConvinced      bool     `json:"be_convinced"`
		BeParalysed      bool     `json:"be_paralysed"`
		BeSummoned       bool     `json:"be_summoned"`
		Behaviour        string   `json:"behaviour"`
		ConvincedMana    int      `json:"convinced_mana"`
		Description      string   `json:"description"`
		ExperiencePoints int      `json:"experience_points"`
		Featured         bool     `json:"featured"`
		Healed           []string `json:"healed"`
		Hitpoints        int      `json:"hitpoints"`
		ImageURL         string   `json:"image_url"`
		Immune           []string `json:"immune"`
		IsLootable       bool     `json:"is_lootable"`
		LootList         []string `json:"loot_list"`
		Name             string   `json:"name"`
		Race             string   `json:"race"`
		SeeInvisible     bool     `json:"see_invisible"`
		Strong           []string `json:"strong"`
		SummonedMana     int      `json:"summoned_mana"`
		Weakness         []string `json:"weakness"`
	}
)
