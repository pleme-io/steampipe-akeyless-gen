package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessOidcApp(t *testing.T) {
	table := tableAkeylessAkeylessOidcApp()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessOidcApp() returned nil")
	}
	if table.Name != "akeyless_akeyless_oidc_app" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_oidc_app", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
