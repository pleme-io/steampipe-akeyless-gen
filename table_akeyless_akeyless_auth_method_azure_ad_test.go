package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodAzureAd(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodAzureAd()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodAzureAd() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_azure_ad" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_azure_ad", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
