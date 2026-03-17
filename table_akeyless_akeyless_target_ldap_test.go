package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetLdap(t *testing.T) {
	table := tableAkeylessAkeylessTargetLdap()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetLdap() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_ldap" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_ldap", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
