package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessGatewayK8sAuthConfig(t *testing.T) {
	table := tableAkeylessAkeylessGatewayK8sAuthConfig()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessGatewayK8sAuthConfig() returned nil")
	}
	if table.Name != "akeyless_akeyless_gateway_k8s_auth_config" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_gateway_k8s_auth_config", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
