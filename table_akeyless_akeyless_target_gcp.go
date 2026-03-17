package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGcp() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_gcp",
		Description: "Manages a GCP target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGcp,
		},
		Columns: akeylessAkeylessTargetGcpColumns(),
	}
}

func akeylessAkeylessTargetGcpColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "gcp_key",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account key JSON (base64)",
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
			Name:        "use_gw_cloud_identity",
			Type:        proto.ColumnType_BOOL,
			Description: "Use gateway cloud identity for authentication",
		},
	}
}

func listAkeylessAkeylessTargetGcp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
