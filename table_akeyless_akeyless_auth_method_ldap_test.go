package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodLdap(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodLdap()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodLdap() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_ldap" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_ldap", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
