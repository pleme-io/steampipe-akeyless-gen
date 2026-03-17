package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretSsh(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretSsh()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretSsh() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_ssh" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_ssh", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
