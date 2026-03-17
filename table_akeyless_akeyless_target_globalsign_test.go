package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGlobalsign(t *testing.T) {
	table := tableAkeylessAkeylessTargetGlobalsign()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGlobalsign() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_globalsign" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_globalsign", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
