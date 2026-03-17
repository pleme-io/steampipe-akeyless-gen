package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGithub(t *testing.T) {
	table := tableAkeylessAkeylessTargetGithub()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGithub() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_github" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_github", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
