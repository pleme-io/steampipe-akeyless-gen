package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEventForwarderServicenow(t *testing.T) {
	table := tableAkeylessAkeylessEventForwarderServicenow()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEventForwarderServicenow() returned nil")
	}
	if table.Name != "akeyless_akeyless_event_forwarder_servicenow" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_event_forwarder_servicenow", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
