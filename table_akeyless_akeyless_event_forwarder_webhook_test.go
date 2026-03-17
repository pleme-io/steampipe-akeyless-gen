package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEventForwarderWebhook(t *testing.T) {
	table := tableAkeylessAkeylessEventForwarderWebhook()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEventForwarderWebhook() returned nil")
	}
	if table.Name != "akeyless_akeyless_event_forwarder_webhook" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_event_forwarder_webhook", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
