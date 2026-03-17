package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGitlab() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_gitlab",
		Description: "Manages a GitLab target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGitlab,
		},
		Columns: akeylessAkeylessTargetGitlabColumns(),
	}
}

func akeylessAkeylessTargetGitlabColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "gitlab_access_token",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab personal access token",
		},
		{
			Name:        "gitlab_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab TLS certificate (PEM)",
		},
		{
			Name:        "gitlab_url",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab URL (for self-hosted)",
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

func listAkeylessAkeylessTargetGitlab(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
