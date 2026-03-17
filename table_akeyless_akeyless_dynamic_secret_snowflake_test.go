package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretSnowflake(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretSnowflake()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretSnowflake() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_snowflake" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_snowflake", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
