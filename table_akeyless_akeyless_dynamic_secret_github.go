package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretGithub() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_github",
		Description: "Manages a GitHub dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretGithub,
		},
		Columns: akeylessAkeylessDynamicSecretGithubColumns(),
	}
}

func akeylessAkeylessDynamicSecretGithubColumns() []*plugin.Column {
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
			Name:        "installation_id",
			Type:        proto.ColumnType_INT,
			Description: "GitHub App installation ID",
		},
		{
			Name:        "installation_organization",
			Type:        proto.ColumnType_STRING,
			Description: "GitHub organization for installation",
		},
		{
			Name:        "installation_repository",
			Type:        proto.ColumnType_STRING,
			Description: "GitHub repository for installation",
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
			Name:        "token_permissions",
			Type:        proto.ColumnType_JSON,
			Description: "GitHub token permissions (JSON)",
		},
		{
			Name:        "token_repositories",
			Type:        proto.ColumnType_JSON,
			Description: "GitHub repositories for token scope (comma-separated)",
		},
		{
			Name:        "token_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "Token TTL",
		},
	}
}

func listAkeylessAkeylessDynamicSecretGithub(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
