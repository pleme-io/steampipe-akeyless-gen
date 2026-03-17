package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretWindows(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretWindows()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretWindows() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_windows" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_windows", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
