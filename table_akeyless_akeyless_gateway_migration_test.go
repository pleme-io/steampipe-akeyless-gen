package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessGatewayMigration(t *testing.T) {
	table := tableAkeylessAkeylessGatewayMigration()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessGatewayMigration() returned nil")
	}
	if table.Name != "akeyless_akeyless_gateway_migration" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_gateway_migration", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
