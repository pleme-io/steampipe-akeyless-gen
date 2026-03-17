package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetArtifactory() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_artifactory",
		Description: "Manages an Artifactory target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetArtifactory,
		},
		Columns: akeylessAkeylessTargetArtifactoryColumns(),
	}
}

func akeylessAkeylessTargetArtifactoryColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "artifactory_admin_name",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory admin username",
		},
		{
			Name:        "artifactory_admin_pwd",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory admin password",
		},
		{
			Name:        "base_url",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory base URL",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key used to encrypt the target secret value (if empty, the\naccount default protectionKey key will be used)",
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

func listAkeylessAkeylessTargetArtifactory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
