package models

type GoogleMaps struct {
	Result GoogleMapsResult `json:"result"`
}

type GoogleMapsPlaceId struct {
	Candidates []GoogleMapsCandidates `json:"candidates"`
}

type GoogleMapsCandidates struct {
	PlaceId string `json:"place_id"`
}

type GoogleMapsResult struct {
	Name         string                 `json:"name"`
	Address      string                 `json:"formatted_address"`
	Phone        string                 `json:"formatted_phone_number"`
	OpeningHours GoogleMapsOpeningHours `json:"opening_hours"`
	Rating       float32                `json:"rating"`
	Reviews      []GoogleMapsReviews    `json:"reviews"`
	Website      string                 `json:"website"`
}

type GoogleMapsOpeningHours struct {
	OpenNow     bool   `json:"open_now"`
	WeekdayText string `json:"weekday_text"`
}

type GoogleMapsReviews struct {
	AuthorName              string  `json:"author_name"`
	AuthorUrl               string  `json:"author_url"`
	ProfilePhotoUrl         string  `json:"profile_photo_url"`
	Rating                  float32 `json:"rating"`
	RelativeTimeDescription string  `json:"relative_time_description"`
	Text                    string  `json:"text"`
}
