package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTargetLetsEncrypt(t *testing.T) {
	table := tableAkeylessAkeylessTargetLetsEncrypt()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTargetLetsEncrypt() returned nil")
	}
	if table.Name != "akeyless_akeyless_target_lets_encrypt" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_target_lets_encrypt", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
