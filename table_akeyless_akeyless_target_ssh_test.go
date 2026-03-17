package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetSsh(t *testing.T) {
	table := tableAkeylessAkeylessTargetSsh()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetSsh() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_ssh" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_ssh", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
