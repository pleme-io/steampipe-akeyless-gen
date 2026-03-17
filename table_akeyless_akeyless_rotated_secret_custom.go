package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRotatedSecretCustom() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_rotated_secret_custom",
		Description: "Manages a custom rotated secret",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRotatedSecretCustom,
		},
		Columns: akeylessAkeylessRotatedSecretCustomColumns(),
	}
}

func akeylessAkeylessRotatedSecretCustomColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "custom_payload",
			Type:        proto.ColumnType_STRING,
			Description: "Secret payload to be sent with rotation request",
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
			Name:        "enable_password_policy",
			Type:        proto.ColumnType_STRING,
			Description: "Enable password policy",
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
			Description: "Length of the password to be generated",
		},
		{
			Name:        "rotate_after_disconnect",
			Type:        proto.ColumnType_STRING,
			Description: "Rotate after SRA session ends [true/false]",
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
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the rotated secret",
		},
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name to associate",
		},
		{
			Name:        "timeout_sec",
			Type:        proto.ColumnType_INT,
			Description: "Maximum allowed time in seconds for the custom rotator to return the results",
		},
		{
			Name:        "use_capital_letters",
			Type:        proto.ColumnType_STRING,
			Description: "Password must contain capital letters [true/false]",
		},
		{
			Name:        "use_lower_letters",
			Type:        proto.ColumnType_STRING,
			Description: "Password must contain lower case letters [true/false]",
		},
		{
			Name:        "use_numbers",
			Type:        proto.ColumnType_STRING,
			Description: "Password must contain numbers [true/false]",
		},
		{
			Name:        "use_special_characters",
			Type:        proto.ColumnType_STRING,
			Description: "Password must contain special characters [true/false]",
		},
	}
}

func listAkeylessAkeylessRotatedSecretCustom(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
