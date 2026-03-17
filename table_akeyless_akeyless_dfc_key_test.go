package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDfcKey(t *testing.T) {
	table := tableAkeylessAkeylessDfcKey()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDfcKey() returned nil")
	}
	if table.Name != "akeyless_akeyless_dfc_key" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dfc_key", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
