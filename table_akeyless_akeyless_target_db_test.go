package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetDb(t *testing.T) {
	table := tableAkeylessAkeylessTargetDb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetDb() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_db" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_db", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
