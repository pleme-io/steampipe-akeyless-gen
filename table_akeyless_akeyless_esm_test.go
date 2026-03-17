package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEsm(t *testing.T) {
	table := tableAkeylessAkeylessEsm()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEsm() returned nil")
	}
	if table.Name != "akeyless_akeyless_esm" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_esm", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
