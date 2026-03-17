package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetAzure(t *testing.T) {
	table := tableAkeylessAkeylessTargetAzure()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetAzure() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_azure" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_azure", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
