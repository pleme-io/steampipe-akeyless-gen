package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAccountCustomField(t *testing.T) {
	table := tableAkeylessAkeylessAccountCustomField()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAccountCustomField() returned nil")
	}
	if table.Name != "akeyless_akeyless_account_custom_field" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_account_custom_field", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
