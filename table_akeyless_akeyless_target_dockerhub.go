package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetDockerhub() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_dockerhub",
		Description: "Manages a Docker Hub target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetDockerhub,
		},
		Columns: akeylessAkeylessTargetDockerhubColumns(),
	}
}

func akeylessAkeylessTargetDockerhubColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "dockerhub_password",
			Type:        proto.ColumnType_STRING,
			Description: "Docker Hub password or access token",
		},
		{
			Name:        "dockerhub_username",
			Type:        proto.ColumnType_STRING,
			Description: "Docker Hub username",
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
	}
}

func listAkeylessAkeylessTargetDockerhub(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
