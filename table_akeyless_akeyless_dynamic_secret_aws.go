package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretAws() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_aws",
		Description: "Manages an AWS dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretAws,
		},
		Columns: akeylessAkeylessDynamicSecretAwsColumns(),
	}
}

func akeylessAkeylessDynamicSecretAwsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_mode",
			Type:        proto.ColumnType_STRING,
			Description: "AWS access mode: iam_user or assume_role",
		},
		{
			Name:        "admin_rotation_interval_days",
			Type:        proto.ColumnType_INT,
			Description: "Admin credentials rotation interval in days",
		},
		{
			Name:        "aws_access_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "AWS access key ID for the producer",
		},
		{
			Name:        "aws_access_secret_key",
			Type:        proto.ColumnType_STRING,
			Description: "AWS secret access key for the producer",
		},
		{
			Name:        "aws_external_id",
			Type:        proto.ColumnType_STRING,
			Description: "The AWS External ID associated with the AWS role (relevant only for assume_role mode)",
		},
		{
			Name:        "aws_role_arns",
			Type:        proto.ColumnType_STRING,
			Description: "AWS role ARNs to assume (comma-separated)",
		},
		{
			Name:        "aws_user_console_access",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable AWS console access for dynamic users",
		},
		{
			Name:        "aws_user_groups",
			Type:        proto.ColumnType_STRING,
			Description: "AWS IAM groups for dynamic users",
		},
		{
			Name:        "aws_user_policies",
			Type:        proto.ColumnType_STRING,
			Description: "AWS IAM policies for dynamic users",
		},
		{
			Name:        "aws_user_programmatic_access",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable programmatic access for dynamic users",
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
			Name:        "enable_admin_rotation",
			Type:        proto.ColumnType_BOOL,
			Description: "Automatic admin credentials rotation",
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
			Name:        "password_length",
			Type:        proto.ColumnType_STRING,
			Description: "The length of the password to be generated",
		},
		{
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Description: "AWS region",
		},
		{
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
		},
		{
			Name:        "session_tags",
			Type:        proto.ColumnType_STRING,
			Description: "String of Key value session tags comma separated, relevant only for Assumed Role",
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
			Name:        "transitive_tag_keys",
			Type:        proto.ColumnType_STRING,
			Description: "String of transitive tag keys space separated, relevant only for Assumed Role",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretAws(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
