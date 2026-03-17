package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretAws(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretAws()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretAws() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_aws" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_aws", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
