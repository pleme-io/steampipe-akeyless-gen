package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGodaddy(t *testing.T) {
	table := tableAkeylessAkeylessTargetGodaddy()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGodaddy() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_godaddy" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_godaddy", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
