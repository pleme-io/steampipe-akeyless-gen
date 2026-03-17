package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetAws(t *testing.T) {
	table := tableAkeylessAkeylessTargetAws()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetAws() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_aws" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_aws", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
