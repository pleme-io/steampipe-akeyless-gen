package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretHanadb(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretHanadb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretHanadb() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_hanadb" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_hanadb", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
