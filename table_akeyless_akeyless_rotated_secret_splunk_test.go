package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretSplunk(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretSplunk()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretSplunk() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_splunk" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_splunk", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
