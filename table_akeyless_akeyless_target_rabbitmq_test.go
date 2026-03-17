package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetRabbitmq(t *testing.T) {
	table := tableAkeylessAkeylessTargetRabbitmq()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetRabbitmq() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_rabbitmq" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_rabbitmq", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
