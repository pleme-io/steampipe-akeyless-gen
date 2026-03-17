package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRoleAuthMethodAssoc(t *testing.T) {
	table := tableAkeylessAkeylessRoleAuthMethodAssoc()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRoleAuthMethodAssoc() returned nil")
	}
	if table.Name != "akeyless_akeyless_role_auth_method_assoc" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_role_auth_method_assoc", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
