package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretGcp(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretGcp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretGcp() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_gcp" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_gcp", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
