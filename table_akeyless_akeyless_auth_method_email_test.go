package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodEmail(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodEmail()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodEmail() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_email" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_email", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
