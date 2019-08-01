package branch

import (
	"regexp"
	"strings"

	"github.com/lapras-inc/yadockeri/config"
)

type Branch struct {
	ID        int
	ProjectID int
	UserID    int
	Name      string
	URL       string
}

func (b *Branch) GetStacName() string {
	reg := regexp.MustCompile(`[/.*@:;~_]`)
	prefix := config.Element("stack_prefix")
	return prefix + strings.ToLower(reg.ReplaceAllString(b.Name, "-"))
}

func (b *Branch) CheckProject(projectID int) bool {
	return b.ProjectID == projectID
}

func (b *Branch) UpdateURL(baseURL string) {
	url := ""
	if strings.HasPrefix(baseURL, "https://") {
		url = strings.Replace(baseURL, "https://", "https://"+b.GetStacName()+".", 1)
	} else if strings.HasPrefix(baseURL, "http://") {
		url = strings.Replace(baseURL, "http://", "http://"+b.GetStacName()+".", 1)
	}
	b.URL = url
}
