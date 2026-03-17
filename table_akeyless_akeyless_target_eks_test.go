package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetEks(t *testing.T) {
	table := tableAkeylessAkeylessTargetEks()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetEks() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_eks" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_eks", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
