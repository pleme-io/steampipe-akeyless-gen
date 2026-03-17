package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodOauth2(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodOauth2()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodOauth2() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_oauth2" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_oauth2", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
