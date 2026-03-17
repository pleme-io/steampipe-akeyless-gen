package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetRabbitmq() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_rabbitmq",
		Description: "Manages a RabbitMQ target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetRabbitmq,
		},
		Columns: akeylessAkeylessTargetRabbitmqColumns(),
	}
}

func akeylessAkeylessTargetRabbitmqColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "max_versions",
			Type:        proto.ColumnType_STRING,
			Description: "Set the maximum number of versions, limited by the account settings defaults.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "rabbitmq_server_password",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ admin password",
		},
		{
			Name:        "rabbitmq_server_uri",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ management API URI",
		},
		{
			Name:        "rabbitmq_server_user",
			Type:        proto.ColumnType_STRING,
			Description: "RabbitMQ admin username",
		},
	}
}

func listAkeylessAkeylessTargetRabbitmq(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
