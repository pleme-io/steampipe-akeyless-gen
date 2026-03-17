package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodOidc(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodOidc()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodOidc() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_oidc" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_oidc", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
