package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretRabbitmq() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_rabbitmq",
		Description: "Manages a RabbitMQ dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretRabbitmq,
		},
		Columns: akeylessAkeylessDynamicSecretRabbitmqColumns(),
	}
}

func akeylessAkeylessDynamicSecretRabbitmqColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "rabbitmq_admin_pwd",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ Admin password",
		},
		{
			Name:        "rabbitmq_admin_user",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ Admin User",
		},
		{
			Name:        "rabbitmq_server_uri",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ management API URI",
		},
		{
			Name:        "rabbitmq_user_conf_permission",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ configure permission for dynamic users",
		},
		{
			Name:        "rabbitmq_user_read_permission",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ read permission for dynamic users",
		},
		{
			Name:        "rabbitmq_user_tags",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ tags for dynamic users (comma-separated)",
		},
		{
			Name:        "rabbitmq_user_vhost",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ vhost for dynamic users",
		},
		{
			Name:        "rabbitmq_user_write_permission",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ write permission for dynamic users",
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

func listAkeylessAkeylessDynamicSecretRabbitmq(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
