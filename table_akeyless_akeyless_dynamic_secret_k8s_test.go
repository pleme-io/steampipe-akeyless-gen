package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretK8s(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretK8s()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretK8s() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_k8s" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_k8s", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
