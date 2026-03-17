package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessRotatedSecretDockerhub(t *testing.T) {
	table := tableAkeylessAkeylessRotatedSecretDockerhub()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessRotatedSecretDockerhub() returned nil")
	}
	if table.Name != "akeyless_akeyless_rotated_secret_dockerhub" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_rotated_secret_dockerhub", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
