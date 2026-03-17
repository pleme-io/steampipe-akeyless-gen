package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretPing(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretPing()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretPing() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_ping" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_ping", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
