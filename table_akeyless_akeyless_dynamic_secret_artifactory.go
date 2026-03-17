package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretArtifactory() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_artifactory",
		Description: "Manages an Artifactory dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretArtifactory,
		},
		Columns: akeylessAkeylessDynamicSecretArtifactoryColumns(),
	}
}

func akeylessAkeylessDynamicSecretArtifactoryColumns() []*plugin.Column {
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
			Name:        "artifactory_token_audience",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory token audience",
		},
		{
			Name:        "artifactory_token_scope",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory token scope",
		},
		{
			Name:        "base_url",
			Type:        proto.ColumnType_STRING,
			Description: "Artifactory base URL",
		},
		{
			Name:        "custom_username_template",
			Type:        proto.ColumnType_STRING,
			Description: "Customize how temporary usernames are generated using go template",
		},
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
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
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

func listAkeylessAkeylessDynamicSecretArtifactory(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
