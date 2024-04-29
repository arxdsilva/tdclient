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

	V4GetSpellResponse struct {
		Information APIInformation           `json:"information"`
		Spell       SpellCompleteInformation `json:"spell"`
	}

	SpellCompleteInformation struct {
		Description         string                   `json:"description"`
		HasRuneInformation  bool                     `json:"has_rune_information"`
		HasSpellInformation bool                     `json:"has_spell_information"`
		ImageURL            string                   `json:"image_url"`
		Name                string                   `json:"name"`
		RuneInformation     RuneInformation          `json:"rune_information"`
		SpellID             string                   `json:"spell_id"`
		SpellInformation    SpellExtendedInformation `json:"spell_information"`
	}

	SpellExtendedInformation struct {
		Amount        int      `json:"amount"`
		City          []string `json:"city"`
		CooldownAlone int      `json:"cooldown_alone"`
		CooldownGroup int      `json:"cooldown_group"`
		DamageType    string   `json:"damage_type"`
		Formula       string   `json:"formula"`
		GroupAttack   bool     `json:"group_attack"`
		GroupHealing  bool     `json:"group_healing"`
		GroupSupport  bool     `json:"group_support"`
		Level         int      `json:"level"`
		Mana          int      `json:"mana"`
		PremiumOnly   bool     `json:"premium_only"`
		Price         int      `json:"price"`
		SoulPoints    int      `json:"soul_points"`
		TypeInstant   bool     `json:"type_instant"`
		TypeRune      bool     `json:"type_rune"`
		Vocation      []string `json:"vocation"`
	}

	RuneInformation struct {
		DamageType   string   `json:"damage_type"`
		GroupAttack  bool     `json:"group_attack"`
		GroupHealing bool     `json:"group_healing"`
		GroupSupport bool     `json:"group_support"`
		Level        int      `json:"level"`
		MagicLevel   int      `json:"magic_level"`
		Vocation     []string `json:"vocation"`
	}
)
