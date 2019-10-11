package values

import (
	"bytes"
	"html/template"
)

// TemplateVariable defines variables to replace template in override values.
type TemplateVariable struct {
	CommitSHA1 string
}

// ReplaceVariables replaces all variables in provided text.
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

// ReplaceVariablesAll replaces all variables in provided text list.
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
