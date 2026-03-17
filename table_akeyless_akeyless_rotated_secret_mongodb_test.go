package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretMongodb(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretMongodb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretMongodb() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_mongodb" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_mongodb", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
