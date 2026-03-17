package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretGitlab(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretGitlab()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretGitlab() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_gitlab" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_gitlab", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
