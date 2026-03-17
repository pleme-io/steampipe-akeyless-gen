package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretSnowflake(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretSnowflake()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretSnowflake() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_snowflake" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_snowflake", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
