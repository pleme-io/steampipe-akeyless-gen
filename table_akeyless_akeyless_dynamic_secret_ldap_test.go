package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretLdap(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretLdap()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretLdap() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_ldap" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_ldap", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
