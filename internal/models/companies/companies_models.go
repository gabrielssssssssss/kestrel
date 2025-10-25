package models

type Company struct {
	Organization OrganizationResult `json:"organization"`
	GoogleMaps   GoogleMaps         `json:"google_maps"`
}
