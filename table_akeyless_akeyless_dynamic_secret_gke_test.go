package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretGke(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretGke()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretGke() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_gke" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_gke", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
