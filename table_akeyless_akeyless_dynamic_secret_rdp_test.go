package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretRdp(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretRdp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretRdp() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_rdp" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_rdp", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
