package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessKmipEnvironment(t *testing.T) {
	table := tableAkeylessAkeylessKmipEnvironment()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessKmipEnvironment() returned nil")
	}
	if table.Name != "akeyless_akeyless_kmip_environment" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_kmip_environment", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
