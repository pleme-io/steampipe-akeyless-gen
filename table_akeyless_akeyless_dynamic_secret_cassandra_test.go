package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretCassandra(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretCassandra()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretCassandra() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_cassandra" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_cassandra", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
