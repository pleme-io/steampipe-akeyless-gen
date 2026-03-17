package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessGroup(t *testing.T) {
	table := tableAkeylessAkeylessGroup()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessGroup() returned nil")
	}
	if table.Name != "akeyless_akeyless_group" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_group", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
