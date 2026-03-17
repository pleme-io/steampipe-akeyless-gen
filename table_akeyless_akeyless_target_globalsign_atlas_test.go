package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetGlobalsignAtlas(t *testing.T) {
	table := tableAkeylessAkeylessTargetGlobalsignAtlas()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetGlobalsignAtlas() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_globalsign_atlas" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_globalsign_atlas", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
