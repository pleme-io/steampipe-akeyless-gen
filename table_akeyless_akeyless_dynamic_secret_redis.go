package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretRedis() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_redis",
		Description: "Manages a Redis dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretRedis,
		},
		Columns: akeylessAkeylessDynamicSecretRedisColumns(),
	}
}

func akeylessAkeylessDynamicSecretRedisColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "acl_rules",
			Type:        proto.ColumnType_STRING,
			Description: "A JSON array list of redis ACL rules to attach to the created user. For available rules see the ACL CAT command https://redis.io/commands/acl-cat\nBy default the user will have permissions to read all keys '[\"~*\", \"+@read\"]'",
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
			Name:        "host",
			Type:        proto.ColumnType_STRING,
			Description: "Redis host",
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
			Name:        "port",
			Type:        proto.ColumnType_STRING,
			Description: "Redis port (default: 6379)",
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

func listAkeylessAkeylessDynamicSecretRedis(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
