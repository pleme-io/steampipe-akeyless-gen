package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodSaml(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodSaml()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodSaml() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_saml" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_saml", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
