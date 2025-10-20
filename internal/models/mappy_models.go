package models

type MappySearch struct {
	POIS []MappySearchResult `json:"pois"`
}

type MappySearchResult struct {
	ID string `json:"id"`
}

type MappyGeo struct {
	LAT           float32                `json:"lat"`
	LNG           float32                `json:"longitude"`
	NAME          string                 `json:"name"`
	PHONE         string                 `json:"phone"`
	RUBRIC        MappyGeoRubric         `json:"rubric"`
	TOWN          string                 `json:"town"`
	TOWNCODE      string                 `json:"townCode"`
	WAYNUMBER     string                 `json:"wayNumber"`
	WAY           string                 `json:"way"`
	STREET        string                 `json:"street"`
	ILLUSTRATION  MappyGeoIllustration   `json:"illustration"`
	MAIL          string                 `json:"mail"`
	URL           string                 `json:"url"`
	COMMUNICATION MappyGeoCommunication  `json:"communication"`
	OPENINGHOURS  string                 `json:"openingHours"`
	OPENINGSTATUS MappyGeoOpenningStatus `json:"openingStatus"`
	REVIEWS       MappyGeoReviews        `json:"reviews"`
	PCODE         string                 `json:"pCode"`
}

type MappyGeoRubric struct {
	ID     string `json:"id"`
	LABEL  string `json:"label"`
	PARENT string `json:"parent"`
}

type MappyGeoIllustration struct {
	URL string `json:"url"`
}

type MappyGeoCommunication struct {
	PHONE   MappyGeoCommunicationPhone `json:"phone"`
	EMAIL   string                     `json:"email"`
	WEBSITE string                     `json:"website"`
}

type MappyGeoCommunicationPhone struct {
	NUMBER string `json:"number"`
}

type MappyGeoOpenningStatus struct {
	STATUS  string `json:"status"`
	MESSAGE string `json:"message"`
}

type MappyGeoReviews struct {
	AVERAGENOTE     float32 `json:"averageNote"`
	MAXNOTE         float32 `json:"maxNote"`
	NUMBEROFREVIEWS int     `json:"numberOfReviews"`
}
