package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessUsc() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_usc",
		Description: "Manages a universal secrets connector (USC) in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessUsc,
		},
		Columns: akeylessAkeylessUscColumns(),
	}
}

func akeylessAkeylessUscColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "binary_value",
			Type:        proto.ColumnType_BOOL,
			Description: "Use this option if the universal secrets value is a base64 encoded binary",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the universal secrets",
		},
		{
			Name:        "namespace",
			Type:        proto.ColumnType_STRING,
			Description: "The namespace (relevant for Hashi vault target)",
		},
		{
			Name:        "object_type",
			Type:        proto.ColumnType_STRING,
			Description: "The object_type field.",
		},
		{
			Name:        "pfx_password",
			Type:        proto.ColumnType_STRING,
			Description: "Optional, the passphrase that protects the private key within the pfx certificate (Relevant only for Azure KV certificates)",
		},
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Description: "Optional, create secret in a specific region (GCP only).\nIf empty, a global secret will be created (provider default).",
		},
		{
			Name:        "secret_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name for the new universal secrets",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the universal secrets",
		},
		{
			Name:        "usc_encryption_key",
			Type:        proto.ColumnType_STRING,
			Description: "Optional, The name of the remote key that used to encrypt the secret value (if empty, the default key will be used)",
		},
		{
			Name:        "usc_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the Universal Secrets Connector item",
		},
		{
			Name:        "value",
			Type:        proto.ColumnType_STRING,
			Description: "Value of the universal secrets item, either text or base64 encoded binary",
		},
	}
}

func listAkeylessAkeylessUsc(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
