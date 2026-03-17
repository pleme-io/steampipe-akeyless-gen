package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretOpenai(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretOpenai()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretOpenai() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_openai" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_openai", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
