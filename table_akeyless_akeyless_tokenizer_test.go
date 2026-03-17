package akeyless

import (
	"testing"
)

func TestTableAkeylessAkeylessTokenizer(t *testing.T) {
	table := tableAkeylessAkeylessTokenizer()
	if table == nil {
		t.Fatal("tableAkeylessAkeylessTokenizer() returned nil")
	}
	if table.Name != "akeyless_akeyless_tokenizer" {
		t.Errorf("expected table name %q, got %q", "akeyless_akeyless_tokenizer", table.Name)
	}
	if len(table.Columns) == 0 {
		t.Error("expected at least one column")
	}
}
