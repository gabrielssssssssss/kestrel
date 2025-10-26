package helpers

import "io/ioutil"

func ReadYaml(path string) string {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(body)
}
