package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	"github.com/gabrielssssssssss/kestrel/pkg/helpers"
	"github.com/gabrielssssssssss/kestrel/pkg/openai"
	search_engine "github.com/gabrielssssssssss/kestrel/pkg/search-engine"
)

type EmployeeStruct struct{}

func NewEmployeeRepository() *EmployeeStruct {
	return &EmployeeStruct{}
}

func (r *EmployeeStruct) FetchEmployee(query string) (models.Employee, error) {
	var payload models.Employee

	rawHtml, err := search_engine.GetRawHtml("linkedin.com", query)
	if err != nil {
		return payload, err
	}

	parsed, err := search_engine.ParseHtml(rawHtml)
	if err != nil {
		return payload, err
	}

	prompt := strings.NewReplacer("[#COMPANY]", query, "[#JSON]", parsed).Replace(helpers.ReadYaml("./prompt/employee.yaml"))
	response, err := openai.PromptTurbo(prompt)
	if err != nil {
		return payload, err
	}

	err = json.Unmarshal([]byte(response), &payload)
	if err != nil {
		return payload, err
	}
	fmt.Println(payload)
	return payload, nil
}
