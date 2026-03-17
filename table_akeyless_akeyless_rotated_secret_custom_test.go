package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretCustom(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretCustom()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretCustom() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_custom" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_custom", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
