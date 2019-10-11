package values

import (
	"bytes"
	"html/template"
)

type TemplateVariable struct {
	CommitSHA1 string
}

func (t *TemplateVariable) ReplaceVariables(text string) (string, error) {
	tmpl, err := template.New("variable").Parse(text)
	if err != nil {
		return "", err
	}
	buffer := &bytes.Buffer{}
	err = tmpl.Execute(buffer, t)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (t *TemplateVariable) ReplaceVariablesAll(text []string) ([]string, error) {
	results := []string{}
	for _, line := range text {
		res, err := t.ReplaceVariables(line)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
