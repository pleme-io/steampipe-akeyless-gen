package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessPkiCertIssuer(t *testing.T) {
	table := tableAkeylessAkeylessPkiCertIssuer()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessPkiCertIssuer() returned nil")
	}
	if table.Name != "akeyless_akeyless_pki_cert_issuer" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_pki_cert_issuer", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
