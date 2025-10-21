package models

type Identity struct {
	FULLNAME   string     `json:"fullname"`
	FIRSTNAME  string     `json:"firstname"`
	LASTNAME   string     `json:"lastname"`
	EMAIL      string     `json:"email"`
	MEDIAS     Medias     `json:"social_medias"`
	RELATIONAL Relational `json:"relational"`
}

type Medias struct {
	YOUTUBE  []string `json:"youtube"`
	TIWTTER  []string `json:"twitter"`
	LINKEDIN []string `json:"linkedin"`
	FACEBOOK []string `json:"facebook"`
	LINKTREE []string `json:"linktree"`
	TIKTOK   []string `json:"tiktok"`
}

type Relational struct {
	RANK        string `json:"rank"`
	TITLE       string `json:"title"`
	URL         string `json:"url"`
	DESCRIPTION string `json:"description"`
}
