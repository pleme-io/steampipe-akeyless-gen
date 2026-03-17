package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEventForwarderEmail(t *testing.T) {
	table := tableAkeylessAkeylessEventForwarderEmail()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEventForwarderEmail() returned nil")
	}
	if table.Name != "akeyless_akeyless_event_forwarder_email" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_event_forwarder_email", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
