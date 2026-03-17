package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretOracle() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_oracle",
		Description: "Manages an Oracle dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretOracle,
		},
		Columns: akeylessAkeylessDynamicSecretOracleColumns(),
	}
}

func akeylessAkeylessDynamicSecretOracleColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "custom_username_template",
			Type:        proto.ColumnType_STRING,
			Description: "Customize how temporary usernames are generated using go template",
		},
		{
			Name:        "db_server_certificates",
			Type:        proto.ColumnType_STRING,
			Description: "Database server certificates (PEM)",
		},
		{
			Name:        "db_server_name",
			Type:        proto.ColumnType_STRING,
			Description: "Database server name for TLS verification",
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
			Name:        "oracle_host",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle host",
		},
		{
			Name:        "oracle_password",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle admin password",
		},
		{
			Name:        "oracle_port",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle port (default: 1521)",
		},
		{
			Name:        "oracle_revocation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle revocation statements",
		},
		{
			Name:        "oracle_screation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle Creation statements",
		},
		{
			Name:        "oracle_service_name",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle service name",
		},
		{
			Name:        "oracle_username",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle admin username",
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

func listAkeylessAkeylessDynamicSecretOracle(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
