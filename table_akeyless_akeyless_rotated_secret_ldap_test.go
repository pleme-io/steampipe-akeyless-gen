package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretLdap(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretLdap()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretLdap() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_ldap" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_ldap", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
