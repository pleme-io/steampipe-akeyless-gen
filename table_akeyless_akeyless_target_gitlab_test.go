package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGitlab(t *testing.T) {
	table := tableAkeylessAkeylessTargetGitlab()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGitlab() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_gitlab" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_gitlab", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
