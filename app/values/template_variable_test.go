package values

import (
	"testing"
)

func TestReplaceVariables(t *testing.T) {
	revision := "test_git_revision"
	v := &TemplateVariable{
		CommitSHA1: revision,
	}

	embeddedText := "image.tag={{.CommitSHA1}}"
	replaced, err := v.ReplaceVariables(embeddedText)
	if err != nil {
		t.Error(err)
	}
	if replaced != "image.tag=test_git_revision" {
		t.Errorf("replaced variable does not match: %s", replaced)
	}

	rawText := "smtp_user=user_name"
	replaced, err = v.ReplaceVariables(rawText)
	if err != nil {
		t.Error(err)
	}
	if replaced != rawText {
		t.Errorf("could not replace raw text: %s", replaced)
	}
}

func TestReplaceVariablesAll(t *testing.T) {
	revision := "test_git_revision"
	v := &TemplateVariable{
		CommitSHA1: revision,
	}

	multipleText := []string{
		"image.tag={{.CommitSHA1}}",
		"smtp_user=hoge",
		"smtp_password=fuga",
	}
	replaced, err := v.ReplaceVariablesAll(multipleText)
	if err != nil {
		t.Error(err)
	}
	if len(replaced) != 3 {
		t.Error("replaced length is not match")
	}

	if replaced[0] != "image.tag=test_git_revision" {
		t.Errorf("replaced variable does not match: %s", replaced[0])
	}

	if replaced[1] != "smtp_user=hoge" {
		t.Errorf("replace variable does not match: %s", replaced[1])
	}

	if replaced[2] != "smtp_password=fuga" {
		t.Errorf("replace variable does not match: %s", replaced[2])
	}
}
