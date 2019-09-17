package branch

import (
	"testing"

	"github.com/h3poteto/yadockeri/config"
)

func TestGetStackName(t *testing.T) {
	normalBranch := &Branch{
		ID:        1,
		ProjectID: 1,
		UserID:    1,
		Name:      "65/feature/add_new",
		URL:       "",
	}

	normalStack := normalBranch.GetStacName()
	if normalStack != config.Element("stack_prefix")+"65-feature-add-new" {
		t.Errorf("Stack name is invalid for URL: %s", normalStack)
	}

	longBranch := &Branch{
		ID:        2,
		ProjectID: 2,
		UserID:    2,
		Name:      "65/feature/xxxxxx-xxxxxx_xxxxxx-xxxxxxx-xxxx-xxxx-xxxxxxxxxxxx",
	}
	longStack := longBranch.GetStacName()
	if longStack != "iss-65-feature-xxxxxx-xxxxxx-xxxxxx-xxxxxxx-xxxx-xxxx" {
		t.Errorf("Stack name is invalid for URL: %s", longStack)
	}
}
