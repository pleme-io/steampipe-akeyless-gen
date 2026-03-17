package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGemini(t *testing.T) {
	table := tableAkeylessAkeylessTargetGemini()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGemini() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_gemini" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_gemini", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
