package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTokenizer() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_tokenizer",
		Description: "Manages a tokenizer in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTokenizer,
		},
		Columns: akeylessAkeylessTokenizerColumns(),
	}
}

func akeylessAkeylessTokenizerColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "alphabet",
			Type:        proto.ColumnType_STRING,
			Description: "Alphabet to use in regexp vaultless tokenization",
		},
		{
			Name:        "decoding_template",
			Type:        proto.ColumnType_STRING,
			Description: "The Decoding output template to use in regexp vaultless tokenization",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_STRING,
			Description: "Protection from accidental deletion of this object [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "encoding_template",
			Type:        proto.ColumnType_STRING,
			Description: "The Encoding output template to use in regexp vaultless tokenization",
		},
		{
			Name:        "encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "AES key name to use in vaultless tokenization",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Tokenizer name",
		},
		{
			Name:        "pattern",
			Type:        proto.ColumnType_STRING,
			Description: "Pattern to use in regexp vaultless tokenization",
		},
		{
			Name:        "tag",
			Type:        proto.ColumnType_JSON,
			Description: "List of the tags attached to this key",
		},
		{
			Name:        "template_type",
			Type:        proto.ColumnType_STRING,
			Description: "Which template type this tokenizer is used for [SSN,CreditCard,USPhoneNumber,Email,Regexp]",
		},
		{
			Name:        "tokenizer_type",
			Type:        proto.ColumnType_STRING,
			Description: "Tokenizer type",
		},
		{
			Name:        "tweak_type",
			Type:        proto.ColumnType_STRING,
			Description: "The tweak type to use in vaultless tokenization [Supplied, Generated, Internal, Masking]",
		},
	}
}

func listAkeylessAkeylessTokenizer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
