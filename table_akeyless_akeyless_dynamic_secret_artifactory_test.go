package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessDynamicSecretArtifactory(t *testing.T) {
	table := tableAkeylessAkeylessDynamicSecretArtifactory()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessDynamicSecretArtifactory() returned nil")
	}
	if table.Name != "akeyless_akeyless_dynamic_secret_artifactory" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_dynamic_secret_artifactory", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
