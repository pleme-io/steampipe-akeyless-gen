package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretHanadb() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_hanadb",
		Description: "Manages a SAP HANA dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretHanadb,
		},
		Columns: akeylessAkeylessDynamicSecretHanadbColumns(),
	}
}

func akeylessAkeylessDynamicSecretHanadbColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "custom_username_template",
			Type:        proto.ColumnType_STRING,
			Description: "Customize how temporary usernames are generated using go template",
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
			Name:        "hana_dbname",
			Type:        proto.ColumnType_STRING,
			Description: "HanaDb Name",
		},
		{
			Name:        "hanadb_create_statements",
			Type:        proto.ColumnType_STRING,
			Description: "HanaDb Creation statements",
		},
		{
			Name:        "hanadb_host",
			Type:        proto.ColumnType_STRING,
			Description: "HANA host",
		},
		{
			Name:        "hanadb_password",
			Type:        proto.ColumnType_STRING,
			Description: "HANA admin password",
		},
		{
			Name:        "hanadb_port",
			Type:        proto.ColumnType_STRING,
			Description: "HANA port (default: 443)",
		},
		{
			Name:        "hanadb_revocation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "HANA SQL revocation statements",
		},
		{
			Name:        "hanadb_username",
			Type:        proto.ColumnType_STRING,
			Description: "HANA admin username",
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
			Name:        "password_length",
			Type:        proto.ColumnType_STRING,
			Description: "The length of the password to be generated",
		},
		{
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the producer",
		},
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name associated with this producer",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretHanadb(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
