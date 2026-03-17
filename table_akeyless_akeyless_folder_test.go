package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessFolder(t *testing.T) {
	table := tableAkeylessAkeylessFolder()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessFolder() returned nil")
	}
	if table.Name != "akeyless_akeyless_folder" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_folder", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
