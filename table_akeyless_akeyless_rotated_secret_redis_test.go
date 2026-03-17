package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretRedis(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretRedis()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretRedis() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_redis" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_redis", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
