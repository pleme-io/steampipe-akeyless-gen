package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodApiKey(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodApiKey()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodApiKey() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_api_key" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_api_key", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
