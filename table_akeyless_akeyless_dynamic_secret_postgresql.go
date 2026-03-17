package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretPostgresql() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_postgresql",
		Description: "Manages a PostgreSQL dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretPostgresql,
		},
		Columns: akeylessAkeylessDynamicSecretPostgresqlColumns(),
	}
}

func akeylessAkeylessDynamicSecretPostgresqlColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "creation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL creation statements",
		},
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
			Name:        "postgresql_db_name",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL DB Name",
		},
		{
			Name:        "postgresql_host",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL host",
		},
		{
			Name:        "postgresql_password",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL admin password",
		},
		{
			Name:        "postgresql_port",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL port (default: 5432)",
		},
		{
			Name:        "postgresql_username",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL admin username",
		},
		{
			Name:        "producer_encryption_key",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "revocation_statement",
			Type:        proto.ColumnType_STRING,
			Description: "PostgreSQL Revocation statements",
		},
		{
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
		},
		{
			Name:        "ssl",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable SSL connection",
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

func listAkeylessAkeylessDynamicSecretPostgresql(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
