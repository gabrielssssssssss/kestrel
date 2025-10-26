package models

type Employee struct {
	Company       string   `json:"company"`
	EmployeeCount int      `json:"employee_count"`
	Worker        []Worker `json:"employees"`
}

type Worker struct {
	Name      string `json:"name"`
	Position  string `json:"position"`
	Hierarchy string `json:"hierarchy_level"`
	Linkedin  string `json:"linkedin"`
	Employed  bool   `json:"employed"`
}
