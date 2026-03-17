package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretCassandra() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_cassandra",
		Description: "Manages a Cassandra dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretCassandra,
		},
		Columns: akeylessAkeylessDynamicSecretCassandraColumns(),
	}
}

func akeylessAkeylessDynamicSecretCassandraColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "cassandra_creation_statements",
			Type:        proto.ColumnType_STRING,
			Description: "Cassandra creation statements (CQL)",
		},
		{
			Name:        "cassandra_hosts",
			Type:        proto.ColumnType_STRING,
			Description: "Cassandra hosts (comma-separated)",
		},
		{
			Name:        "cassandra_password",
			Type:        proto.ColumnType_STRING,
			Description: "Cassandra admin password",
		},
		{
			Name:        "cassandra_port",
			Type:        proto.ColumnType_STRING,
			Description: "Cassandra port (default: 9042)",
		},
		{
			Name:        "cassandra_username",
			Type:        proto.ColumnType_STRING,
			Description: "Cassandra admin username",
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
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
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

func listAkeylessAkeylessDynamicSecretCassandra(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
