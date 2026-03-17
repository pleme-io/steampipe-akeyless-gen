package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretMysql(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretMysql()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretMysql() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_mysql" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_mysql", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
