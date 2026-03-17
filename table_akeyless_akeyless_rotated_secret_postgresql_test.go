package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretPostgresql(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretPostgresql()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretPostgresql() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_postgresql" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_postgresql", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
