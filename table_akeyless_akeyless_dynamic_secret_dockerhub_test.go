package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretDockerhub(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretDockerhub()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretDockerhub() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_dockerhub" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_dockerhub", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
