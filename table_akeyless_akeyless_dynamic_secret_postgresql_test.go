package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretPostgresql(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretPostgresql()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretPostgresql() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_postgresql" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_postgresql", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
