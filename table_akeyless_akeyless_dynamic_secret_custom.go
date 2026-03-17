package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretCustom() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_custom",
		Description: "Manages a custom dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretCustom,
		},
		Columns: akeylessAkeylessDynamicSecretCustomColumns(),
	}
}

func akeylessAkeylessDynamicSecretCustomColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "admin_rotation_interval_days",
			Type:        proto.ColumnType_INT,
			Description: "Admin credentials rotation interval in days",
		},
		{
			Name:        "create_sync_url",
			Type:        proto.ColumnType_STRING,
			Description: "URL to call on create/sync",
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
			Name:        "enable_admin_rotation",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable admin credentials rotation",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic secret name",
		},
		{
			Name:        "payload",
			Type:        proto.ColumnType_STRING,
			Description: "Custom payload to send to the URL",
		},
		{
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "revoke_sync_url",
			Type:        proto.ColumnType_STRING,
			Description: "URL to call on revoke",
		},
		{
			Name:        "rotate_sync_url",
			Type:        proto.ColumnType_STRING,
			Description: "URL to call on rotate",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the producer",
		},
		{
			Name:        "timeout_sec",
			Type:        proto.ColumnType_INT,
			Description: "Timeout in seconds for custom calls",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretCustom(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
