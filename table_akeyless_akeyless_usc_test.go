package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessUsc(t *testing.T) {
	table := tableAkeylessAkeylessUsc()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessUsc() returned nil")
	}
	if table.Name != "akeyless_akeyless_usc" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_usc", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
