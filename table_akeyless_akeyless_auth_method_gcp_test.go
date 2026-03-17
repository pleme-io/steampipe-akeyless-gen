package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodGcp(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodGcp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodGcp() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_gcp" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_gcp", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
