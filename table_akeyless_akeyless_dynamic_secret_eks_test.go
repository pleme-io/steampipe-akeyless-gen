package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretEks(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretEks()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretEks() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_eks" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_eks", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
