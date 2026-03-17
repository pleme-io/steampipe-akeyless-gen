package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretHanadb(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretHanadb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretHanadb() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_hanadb" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_hanadb", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
