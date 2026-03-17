package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretOpenai(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretOpenai()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretOpenai() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_openai" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_openai", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
