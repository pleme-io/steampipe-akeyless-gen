package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGke(t *testing.T) {
	table := tableAkeylessAkeylessTargetGke()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGke() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_gke" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_gke", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
