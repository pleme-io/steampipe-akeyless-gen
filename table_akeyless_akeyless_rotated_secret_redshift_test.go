package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretRedshift(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretRedshift()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretRedshift() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_redshift" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_redshift", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
