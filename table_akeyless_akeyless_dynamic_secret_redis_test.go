package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretRedis(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretRedis()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretRedis() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_redis" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_redis", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
