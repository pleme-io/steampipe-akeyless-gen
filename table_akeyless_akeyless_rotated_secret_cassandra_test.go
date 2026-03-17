package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretCassandra(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretCassandra()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretCassandra() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_cassandra" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_cassandra", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
