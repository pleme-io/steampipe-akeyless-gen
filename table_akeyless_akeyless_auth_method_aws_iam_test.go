package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessAuthMethodAwsIam(t *testing.T) {
	table := tableAkeylessAkeylessAuthMethodAwsIam()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessAuthMethodAwsIam() returned nil")
	}
	if table.Name != "akeyless_akeyless_auth_method_aws_iam" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_auth_method_aws_iam", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
