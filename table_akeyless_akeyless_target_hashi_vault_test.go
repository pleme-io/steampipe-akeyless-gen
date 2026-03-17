package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetHashiVault(t *testing.T) {
	table := tableAkeylessAkeylessTargetHashiVault()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetHashiVault() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_hashi_vault" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_hashi_vault", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
