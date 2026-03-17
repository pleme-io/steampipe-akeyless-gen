package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretGithub(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretGithub()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretGithub() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_github" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_github", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
