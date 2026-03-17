package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretMysql() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_mysql",
		Description: "Manages a MySQL dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretMysql,
		},
		Columns: akeylessAkeylessDynamicSecretMysqlColumns(),
	}
}

func akeylessAkeylessDynamicSecretMysqlColumns() []*plugin.Column {
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
			Name:        "mysql_dbname",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL database name",
		},
		{
			Name:        "mysql_host",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL host",
		},
		{
			Name:        "mysql_password",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL admin password",
		},
		{
			Name:        "mysql_port",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL port (default: 3306)",
		},
		{
			Name:        "mysql_revocation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL revocation statements",
		},
		{
			Name:        "mysql_screation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL Creation statements",
		},
		{
			Name:        "mysql_username",
			Type:        proto.ColumnType_STRING,
			Description: "MySQL admin username",
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
			Name:        "ssl",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable SSL connection",
		},
		{
			Name:        "ssl_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "SSL certificate (PEM)",
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

func listAkeylessAkeylessDynamicSecretMysql(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
