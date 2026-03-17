package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodOci(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodOci()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodOci() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_oci" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_oci", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
