package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetPing(t *testing.T) {
	table := tableAkeylessAkeylessTargetPing()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetPing() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_ping" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_ping", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
