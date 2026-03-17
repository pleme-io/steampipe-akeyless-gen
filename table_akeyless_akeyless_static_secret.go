package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessStaticSecret() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_static_secret",
		Description: "Manages a static secret in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessStaticSecret,
		},
		Columns: akeylessAkeylessStaticSecretColumns(),
	}
}

func akeylessAkeylessStaticSecretColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "accessibility",
			Type:        proto.ColumnType_STRING,
			Description: "Secret accessibility: regular or personal",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "format",
			Type:        proto.ColumnType_STRING,
			Description: "Secret format [text/json/key-value] (relevant only for type 'generic')",
		},
		{
			Name:        "max_versions",
			Type:        proto.ColumnType_STRING,
			Description: "Set the maximum number of versions, limited by the account settings defaults.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Secret name",
		},
		{
			Name:        "protection_key",
			Type:        proto.ColumnType_STRING,
			Description: "Encryption key name for the secret",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Add tags attached to this object",
		},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "The secret sub type [generic/password]",
		},
		{
			Name:        "value",
			Type:        proto.ColumnType_STRING,
			Description: "The secret value",
		},
	}
}

func listAkeylessAkeylessStaticSecret(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
