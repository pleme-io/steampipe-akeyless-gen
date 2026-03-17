package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetSplunk(t *testing.T) {
	table := tableAkeylessAkeylessTargetSplunk()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetSplunk() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_splunk" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_splunk", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
