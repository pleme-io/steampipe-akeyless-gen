package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodCert(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodCert()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodCert() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_cert" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_cert", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
