package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessGatewayAllowedAccess(t *testing.T) {
	table := tableAkeylessAkeylessGatewayAllowedAccess()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessGatewayAllowedAccess() returned nil")
	}
	if table.Name != "akeyless_akeyless_gateway_allowed_access" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_gateway_allowed_access", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
