package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetSectigo(t *testing.T) {
	table := tableAkeylessAkeylessTargetSectigo()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetSectigo() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_sectigo" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_sectigo", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
