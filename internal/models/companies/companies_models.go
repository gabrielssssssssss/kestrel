package models

type Company struct {
	Organization OrganizationResult `json:"organization"`
	Maps         Maps               `json:"google_maps"`
}
