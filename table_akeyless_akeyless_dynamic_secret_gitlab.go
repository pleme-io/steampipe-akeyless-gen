package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretGitlab() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_gitlab",
		Description: "Manages a GitLab dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretGitlab,
		},
		Columns: akeylessAkeylessDynamicSecretGitlabColumns(),
	}
}

func akeylessAkeylessDynamicSecretGitlabColumns() []*plugin.Column {
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
			Name:        "gitlab_access_token",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab personal access token",
		},
		{
			Name:        "gitlab_access_type",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab access type: personal or project",
		},
		{
			Name:        "gitlab_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab TLS certificate (PEM)",
		},
		{
			Name:        "gitlab_role",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab role for dynamic tokens",
		},
		{
			Name:        "gitlab_token_scopes",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab token scopes (comma-separated)",
		},
		{
			Name:        "gitlab_url",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab URL (for self-hosted)",
		},
		{
			Name:        "group_name",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab group name",
		},
		{
			Name:        "installation_organization",
			Type:        proto.ColumnType_STRING,
			Description: "GitLab organization",
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
			Name:        "ttl",
			Type:        proto.ColumnType_STRING,
			Description: "Access Token TTL",
		},
	}
}

func listAkeylessAkeylessDynamicSecretGitlab(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
