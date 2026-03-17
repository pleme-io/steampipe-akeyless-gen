package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessPolicy(t *testing.T) {
	table := tableAkeylessAkeylessPolicy()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessPolicy() returned nil")
	}
	if table.Name != "akeyless_akeyless_policy" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_policy", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
