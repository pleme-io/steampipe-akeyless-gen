package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetPing() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_ping",
		Description: "Manages a Ping Identity target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetPing,
		},
		Columns: akeylessAkeylessTargetPingColumns(),
	}
}

func akeylessAkeylessTargetPingColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "administrative_port",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity administrative port",
		},
		{
			Name:        "authorization_port",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity authorization port",
		},
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
			Name:        "ping_url",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity server URL",
		},
		{
			Name:        "privileged_user",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity privileged username",
		},
	}
}

func listAkeylessAkeylessTargetPing(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
