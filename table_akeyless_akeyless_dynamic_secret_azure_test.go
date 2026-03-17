package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretAzure(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretAzure()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretAzure() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_azure" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_azure", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
