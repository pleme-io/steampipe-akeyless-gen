package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretRedshift(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretRedshift()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretRedshift() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_redshift" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_redshift", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
