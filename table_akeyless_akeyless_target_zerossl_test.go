package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetZerossl(t *testing.T) {
	table := tableAkeylessAkeylessTargetZerossl()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetZerossl() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_zerossl" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_zerossl", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
