package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretVenafi(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretVenafi()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretVenafi() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_venafi" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_venafi", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
