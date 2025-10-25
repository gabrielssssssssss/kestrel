package models

type Company struct {
	Organization Organization `json:"organization"`
	GoogleMaps   GoogleMaps   `json:"google_maps"`
}
