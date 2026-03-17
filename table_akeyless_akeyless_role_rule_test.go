package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRoleRule(t *testing.T) {
	table := tableAkeylessAkeylessRoleRule()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRoleRule() returned nil")
	}
	if table.Name != "akeyless_akeyless_role_rule" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_role_rule", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
