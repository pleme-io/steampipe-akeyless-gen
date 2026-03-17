package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretOracle(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretOracle()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretOracle() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_oracle" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_oracle", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
