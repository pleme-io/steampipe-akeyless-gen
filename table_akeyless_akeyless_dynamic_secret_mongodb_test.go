package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretMongodb(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretMongodb()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretMongodb() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_mongodb" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_mongodb", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
