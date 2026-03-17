package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetLinked(t *testing.T) {
	table := tableAkeylessAkeylessTargetLinked()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetLinked() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_linked" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_linked", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
