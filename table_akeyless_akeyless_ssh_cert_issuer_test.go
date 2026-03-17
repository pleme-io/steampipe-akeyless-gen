package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessSshCertIssuer(t *testing.T) {
	table := tableAkeylessAkeylessSshCertIssuer()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessSshCertIssuer() returned nil")
	}
	if table.Name != "akeyless_akeyless_ssh_cert_issuer" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_ssh_cert_issuer", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
