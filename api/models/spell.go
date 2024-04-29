package models

type (
	V4GetSpellsResponse struct {
		Information APIInformation `json:"information"`
		Spells      SpellsInfo     `json:"spells"`
	}

	SpellsInfo struct {
		SpellList    []SpellInformation `json:"spell_list"`
		SpellsFilter string             `json:"spells_filter"`
	}

	SpellInformation struct {
		Formula      string `json:"formula"`
		GroupAttack  bool   `json:"group_attack"`
		GroupHealing bool   `json:"group_healing"`
		GroupSupport bool   `json:"group_support"`
		Level        int    `json:"level"`
		Mana         int    `json:"mana"`
		Name         string `json:"name"`
		PremiumOnly  bool   `json:"premium_only"`
		Price        int    `json:"price"`
		SpellID      string `json:"spell_id"`
		TypeInstant  bool   `json:"type_instant"`
		TypeRune     bool   `json:"type_rune"`
	}
)
