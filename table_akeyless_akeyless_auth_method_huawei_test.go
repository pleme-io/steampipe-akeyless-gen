package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodHuawei(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodHuawei()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodHuawei() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_huawei" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_huawei", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
