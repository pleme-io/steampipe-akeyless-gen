package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetK8s(t *testing.T) {
	table := tableAkeylessAkeylessTargetK8s()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetK8s() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_k8s" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_k8s", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
