package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretGcp(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretGcp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretGcp() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_gcp" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_gcp", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
