package repository

import (
	"encoding/json"
	"strings"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	search_engine "github.com/gabrielssssssssss/kestrel/pkg/google/cse"
	"github.com/gabrielssssssssss/kestrel/pkg/helpers"
	"github.com/gabrielssssssssss/kestrel/pkg/openai"
)

type EmployeeStruct struct{}

func NewEmployeeRepository() *EmployeeStruct {
	return &EmployeeStruct{}
}

func (r *EmployeeStruct) FetchEmployee(query string, naf string) (models.Employee, error) {
	var payload models.Employee

	rawHtml, err := search_engine.GetRawHtml("fr.linkedin.com/in", query)
	if err != nil {
		return payload, err
	}

	parsed, err := search_engine.ParseHtml(rawHtml)
	if err != nil {
		return payload, err
	}

	prompt := strings.NewReplacer("[#COMPANY]", query, "[#NAF]", naf, "[#JSON]", parsed).Replace(helpers.ReadYaml("./prompt/employee.yaml"))
	response, err := openai.PromptTurbo(prompt)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(response), &payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}
