package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRotatedSecretSplunk() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_rotated_secret_splunk",
		Description: "Manages a Splunk rotated secret",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRotatedSecretSplunk,
		},
		Columns: akeylessAkeylessRotatedSecretSplunkColumns(),
	}
}

func akeylessAkeylessRotatedSecretSplunkColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "Token audience for Splunk token creation (required for rotator-type=token)",
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
			Name:        "expiration_date",
			Type:        proto.ColumnType_STRING,
			Description: "Token expiration date in YYYY-MM-DD format (required for rotator-type=token when manual rotation is selected and no existing token is provided). Time will be set to 00:00 UTC.",
		},
		{
			Name:        "hec_token",
			Type:        proto.ColumnType_STRING,
			Description: "Current Splunk HEC token value to store (relevant only for rotator-type=hec-token). If not provided, a new HEC input will be created in Splunk.",
		},
		{
			Name:        "hec_token_name",
			Type:        proto.ColumnType_STRING,
			Description: "Splunk HEC input name to manage  (required for rotator-type=hec-token)",
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
			Name:        "rotated_password",
			Type:        proto.ColumnType_STRING,
			Description: "Rotated-username password (relevant only for rotator-type=password)",
		},
		{
			Name:        "rotated_username",
			Type:        proto.ColumnType_STRING,
			Description: "Username to be rotated",
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
			Description: "The rotator type: target or password",
		},
		{
			Name:        "splunk_token",
			Type:        proto.ColumnType_STRING,
			Description: "Current Splunk authentication token to store (relevant only for rotator-type=token). If not provided, a new token will be created in Splunk.",
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
			Name:        "token_owner",
			Type:        proto.ColumnType_STRING,
			Description: "Splunk token owner username (relevant only for rotator-type=token)",
		},
	}
}

func listAkeylessAkeylessRotatedSecretSplunk(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
