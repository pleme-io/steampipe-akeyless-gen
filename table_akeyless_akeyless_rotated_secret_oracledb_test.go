package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretOracledb(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretOracledb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretOracledb() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_oracledb" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_oracledb", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
