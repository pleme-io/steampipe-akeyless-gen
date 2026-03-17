package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetWeb(t *testing.T) {
	table := tableAkeylessAkeylessTargetWeb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetWeb() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_web" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_web", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
