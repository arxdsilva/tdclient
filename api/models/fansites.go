package models

type (
	V4GetFansitesResponse struct {
		Fansites struct {
			Promoted    []PromotedFansite  `json:"promoted"`
			Supported   []SupportedFansite `json:"fansites"`
			Information APIInformation     `json:"information"`
		}
	}

	PromotedFansite struct {
		Contact        string      `json:"contact"`
		ContentType    ContentType `json:"content_type"`
		FansiteItem    bool        `json:"fansite_item"`
		FansiteItemURL string      `json:"fansite_item_url"`
		Homepage       string      `json:"homepage"`
		Languages      []string    `json:"languages"`
		LogoURL        string      `json:"logo_url"`
		Name           string      `json:"name"`
		SocialMedia    SocialMedia `json:"social_media"`
		Specials       []string    `json:"specials"`
	}

	SupportedFansite struct {
		Contact        string      `json:"contact"`
		ContentType    ContentType `json:"content_type"`
		FansiteItem    bool        `json:"fansite_item"`
		FansiteItemURL string      `json:"fansite_item_url"`
		Homepage       string      `json:"homepage"`
		Languages      []string    `json:"languages"`
		LogoURL        string      `json:"logo_url"`
		Name           string      `json:"name"`
		SocialMedia    SocialMedia `json:"supported"`
	}

	ContentType struct {
		Statistics bool `json:"statistics"`
		Texts      bool `json:"texts"`
		Tools      bool `json:"tools"`
		Wiki       bool `json:"wiki"`
	}

	SocialMedia struct {
		Discord   bool `json:"discord"`
		Facebook  bool `json:"facebook"`
		Instagram bool `json:"instagram"`
		Reddit    bool `json:"reddit"`
		Twitch    bool `json:"twitch"`
		Twitter   bool `json:"twitter"`
		Youtube   bool `json:"youtube"`
	}
)
