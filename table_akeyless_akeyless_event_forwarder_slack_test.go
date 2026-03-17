package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEventForwarderSlack(t *testing.T) {
	table := tableAkeylessAkeylessEventForwarderSlack()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEventForwarderSlack() returned nil")
	}
	if table.Name != "akeyless_akeyless_event_forwarder_slack" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_event_forwarder_slack", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
