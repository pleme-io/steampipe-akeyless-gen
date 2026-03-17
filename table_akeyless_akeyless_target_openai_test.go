package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetOpenai(t *testing.T) {
	table := tableAkeylessAkeylessTargetOpenai()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetOpenai() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_openai" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_openai", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
