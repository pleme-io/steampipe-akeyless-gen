package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRotatedSecretOpenai() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_rotated_secret_openai",
		Description: "Manages an OpenAI rotated secret",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRotatedSecretOpenai,
		},
		Columns: akeylessAkeylessRotatedSecretOpenaiColumns(),
	}
}

func akeylessAkeylessRotatedSecretOpenaiColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "API key to rotate (relevant only for rotator-type=api-key)",
		},
		{
			Name:        "api_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "Admin API key ID to rotate (relevant only for rotator-type=api-key)",
		},
		{
			Name:        "authentication_credentials",
			Type:        proto.ColumnType_STRING,
			Description: "Credentials to connect with: use-user-creds or use-target-creds",
		},
		{
			Name:        "auto_rotate",
			Type:        proto.ColumnType_STRING,
			Description: "Whether to automatically rotate every rotation-interval days [true/false]",
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
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "Encryption key name for the secret value",
		},
		{
			Name:        "max_versions",
			Type:        proto.ColumnType_STRING,
			Description: "Maximum number of versions",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Rotated secret name",
		},
		{
			Name:        "password_length",
			Type:        proto.ColumnType_STRING,
			Description: "The length of the password to be generated",
		},
		{
			Name:        "rotation_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the rotation of the item would you like to be notified",
		},
		{
			Name:        "rotation_hour",
			Type:        proto.ColumnType_INT,
			Description: "The hour of the rotation in UTC",
		},
		{
			Name:        "rotation_interval",
			Type:        proto.ColumnType_STRING,
			Description: "Days between every automatic key rotation (1-365)",
		},
		{
			Name:        "rotator_type",
			Type:        proto.ColumnType_STRING,
			Description: "The rotator type: target or api-key",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the rotated secret",
		},
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name to associate",
		},
	}
}

func listAkeylessAkeylessRotatedSecretOpenai(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
