package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGcp(t *testing.T) {
	table := tableAkeylessAkeylessTargetGcp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGcp() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_gcp" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_gcp", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
