package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetWindows(t *testing.T) {
	table := tableAkeylessAkeylessTargetWindows()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetWindows() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_windows" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_windows", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
