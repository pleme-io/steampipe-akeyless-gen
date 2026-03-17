package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessCertificate(t *testing.T) {
	table := tableAkeylessAkeylessCertificate()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessCertificate() returned nil")
	}
	if table.Name != "akeyless_akeyless_certificate" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_certificate", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
