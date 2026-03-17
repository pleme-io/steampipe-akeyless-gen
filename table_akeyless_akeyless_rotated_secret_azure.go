package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRotatedSecretAzure() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_rotated_secret_azure",
		Description: "Manages an Azure rotated secret",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRotatedSecretAzure,
		},
		Columns: akeylessAkeylessRotatedSecretAzureColumns(),
	}
}

func akeylessAkeylessRotatedSecretAzureColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_id",
			Type:        proto.ColumnType_STRING,
			Description: "API ID to rotate (relevant only for rotator-type=api-key)",
		},
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "API key to rotate (relevant only for rotator-type=api-key)",
		},
		{
			Name:        "application_id",
			Type:        proto.ColumnType_STRING,
			Description: "Id of the azure app that hold the serect to be rotated (relevant only for rotator-type=api-key & authentication-credentials=use-target-creds)",
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
			Name:        "explicitly_set_sa",
			Type:        proto.ColumnType_STRING,
			Description: "If set, explicitly provide the storage account details [true/false]",
		},
		{
			Name:        "grace_rotation",
			Type:        proto.ColumnType_STRING,
			Description: "Enable graceful rotation (keep both versions temporarily). When enabled, a new secret version is created while the previous version is kept for the grace period, so both versions exist for a limited time. [true/false]",
		},
		{
			Name:        "grace_rotation_hour",
			Type:        proto.ColumnType_INT,
			Description: "The Hour of the grace rotation in UTC",
		},
		{
			Name:        "grace_rotation_interval",
			Type:        proto.ColumnType_STRING,
			Description: "The number of days to wait before deleting the old key (must be bigger than rotation-interval)",
		},
		{
			Name:        "grace_rotation_timing",
			Type:        proto.ColumnType_STRING,
			Description: "When to create the new version relative to the rotation date [after/before]",
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
			Name:        "resource_group_name",
			Type:        proto.ColumnType_STRING,
			Description: "The resource group name (only relevant when explicitly-set-sa=true)",
		},
		{
			Name:        "resource_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the storage account (only relevant when explicitly-set-sa=true)",
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
			Name:        "rotator_type",
			Type:        proto.ColumnType_STRING,
			Description: "The rotator type: target or api-key",
		},
		{
			Name:        "storage_account_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the storage account key to rotate [key1/key2/kerb1/kerb2] (relevat to azure-storage-account)",
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

func listAkeylessAkeylessRotatedSecretAzure(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
