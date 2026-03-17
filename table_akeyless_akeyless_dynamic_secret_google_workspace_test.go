package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretGoogleWorkspace(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretGoogleWorkspace()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretGoogleWorkspace() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_google_workspace" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_google_workspace", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
