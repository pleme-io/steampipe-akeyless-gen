package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessStaticSecret(t *testing.T) {
	table := tableAkeylessAkeylessStaticSecret()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessStaticSecret() returned nil")
	}
	if table.Name != "akeyless_akeyless_static_secret" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_static_secret", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
