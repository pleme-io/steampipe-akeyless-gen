package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessEventForwarderTeams(t *testing.T) {
	table := tableAkeylessAkeylessEventForwarderTeams()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessEventForwarderTeams() returned nil")
	}
	if table.Name != "akeyless_akeyless_event_forwarder_teams" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_event_forwarder_teams", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
