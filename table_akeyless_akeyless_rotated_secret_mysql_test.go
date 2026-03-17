package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretMysql(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretMysql()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretMysql() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_mysql" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_mysql", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
