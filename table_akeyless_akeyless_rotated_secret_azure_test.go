package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretAzure(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretAzure()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretAzure() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_azure" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_azure", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
