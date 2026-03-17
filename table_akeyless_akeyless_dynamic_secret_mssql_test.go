package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretMssql(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretMssql()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretMssql() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_mssql" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_mssql", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
