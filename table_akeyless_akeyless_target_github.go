package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGithub() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_github",
		Description: "Manages a GitHub target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGithub,
		},
		Columns: akeylessAkeylessTargetGithubColumns(),
	}
}

func akeylessAkeylessTargetGithubColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "github_app_id",
			Type:        proto.ColumnType_INT,
			Description: "GitHub App ID",
		},
		{
			Name:        "github_app_private_key",
			Type:        proto.ColumnType_STRING,
			Description: "GitHub App private key (PEM)",
		},
		{
			Name:        "github_base_url",
			Type:        proto.ColumnType_STRING,
			Description: "GitHub base URL (for GitHub Enterprise)",
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

func listAkeylessAkeylessTargetGithub(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
