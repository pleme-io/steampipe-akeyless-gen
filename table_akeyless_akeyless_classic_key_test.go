package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessClassicKey(t *testing.T) {
	table := tableAkeylessAkeylessClassicKey()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessClassicKey() returned nil")
	}
	if table.Name != "akeyless_akeyless_classic_key" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_classic_key", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
