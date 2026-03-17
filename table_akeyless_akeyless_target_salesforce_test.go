package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetSalesforce(t *testing.T) {
	table := tableAkeylessAkeylessTargetSalesforce()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetSalesforce() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_salesforce" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_salesforce", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
