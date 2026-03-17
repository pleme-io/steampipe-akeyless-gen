package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessKmipClient(t *testing.T) {
	table := tableAkeylessAkeylessKmipClient()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessKmipClient() returned nil")
	}
	if table.Name != "akeyless_akeyless_kmip_client" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_kmip_client", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
