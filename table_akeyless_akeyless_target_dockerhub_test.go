package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetDockerhub(t *testing.T) {
	table := tableAkeylessAkeylessTargetDockerhub()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetDockerhub() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_dockerhub" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_dockerhub", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
