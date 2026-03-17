package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretRabbitmq(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretRabbitmq()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretRabbitmq() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_rabbitmq" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_rabbitmq", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
