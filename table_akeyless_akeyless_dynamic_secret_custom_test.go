package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretCustom(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretCustom()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretCustom() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_custom" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_custom", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
