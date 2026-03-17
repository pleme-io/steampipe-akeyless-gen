package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessPolicy() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_policy",
		Description: "Manages a policy in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessPolicy,
		},
		Columns: akeylessAkeylessPolicyColumns(),
	}
}

func akeylessAkeylessPolicyColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "allowed_algorithms",
			Type:        proto.ColumnType_JSON,
			Description: "Specify allowed key algorithms (e.g., [RSA2048,AES128GCM])",
		},
		{
			Name:        "allowed_key_names",
			Type:        proto.ColumnType_JSON,
			Description: "Specify allowed protection key names. To enforce using the account's default protection key, use 'default-account-key'",
		},
		{
			Name:        "allowed_key_types",
			Type:        proto.ColumnType_JSON,
			Description: "Specify allowed key protection types (dfc, classic-key)",
		},
		{
			Name:        "max_rotation_interval_days",
			Type:        proto.ColumnType_INT,
			Description: "Set the maximum rotation interval for automatic key rotation.",
		},
		{
			Name:        "object_types",
			Type:        proto.ColumnType_JSON,
			Description: "The object types this policy will apply to (items, targets). If not provided, defaults to [items, targets].",
		},
		{
			Name:        "path",
			Type:        proto.ColumnType_STRING,
			Description: "The path the policy refers to",
		},
	}
}

func listAkeylessAkeylessPolicy(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
