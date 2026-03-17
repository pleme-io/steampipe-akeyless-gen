package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRole(t *testing.T) {
	table := tableAkeylessAkeylessRole()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRole() returned nil")
	}
	if table.Name != "akeyless_akeyless_role" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_role", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
