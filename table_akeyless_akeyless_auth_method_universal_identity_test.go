package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodUniversalIdentity(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodUniversalIdentity()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodUniversalIdentity() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_universal_identity" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_universal_identity", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
