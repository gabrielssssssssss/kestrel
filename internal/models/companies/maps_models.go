package models

type Maps struct {
	Result MapsResult `json:"result"`
}

type MapsPlaceId struct {
	Candidates []MapsCandidates `json:"candidates"`
}

type MapsCandidates struct {
	PlaceId string `json:"place_id"`
}

type MapsResult struct {
	Name         string           `json:"name"`
	Address      string           `json:"formatted_address"`
	Phone        string           `json:"formatted_phone_number"`
	OpeningHours MapsOpeningHours `json:"opening_hours"`
	Rating       float32          `json:"rating"`
	Reviews      []MapsReviews    `json:"reviews"`
	Website      string           `json:"website"`
}

type MapsOpeningHours struct {
	OpenNow     bool     `json:"open_now"`
	WeekdayText []string `json:"weekday_text"`
}

type MapsReviews struct {
	AuthorName              string  `json:"author_name"`
	AuthorUrl               string  `json:"author_url"`
	ProfilePhotoUrl         string  `json:"profile_photo_url"`
	Rating                  float32 `json:"rating"`
	RelativeTimeDescription string  `json:"relative_time_description"`
	Text                    string  `json:"text"`
}
