package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodKerberos(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodKerberos()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodKerberos() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_kerberos" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_kerberos", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
