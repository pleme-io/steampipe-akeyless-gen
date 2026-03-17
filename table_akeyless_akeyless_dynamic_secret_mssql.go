package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretMssql() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_mssql",
		Description: "Manages an MSSQL dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretMssql,
		},
		Columns: akeylessAkeylessDynamicSecretMssqlColumns(),
	}
}

func akeylessAkeylessDynamicSecretMssqlColumns() []*plugin.Column {
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
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "mssql_allowed_db_names",
			Type:        proto.ColumnType_STRING,
			Description: "CSV of allowed DB names for runtime selection when getting the secret value.\nEmpty => use target DB only; \"*\" => any DB allowed; One or more names => user must choose from this list",
		},
		{
			Name:        "mssql_create_statements",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL Creation statements",
		},
		{
			Name:        "mssql_dbname",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL database name",
		},
		{
			Name:        "mssql_host",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL host",
		},
		{
			Name:        "mssql_password",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL admin password",
		},
		{
			Name:        "mssql_port",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL port (default: 1433)",
		},
		{
			Name:        "mssql_revocation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL revocation statements",
		},
		{
			Name:        "mssql_username",
			Type:        proto.ColumnType_STRING,
			Description: "MSSQL admin username",
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
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
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

func listAkeylessAkeylessDynamicSecretMssql(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
