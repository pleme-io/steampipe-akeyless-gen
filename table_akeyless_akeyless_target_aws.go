package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetAws() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_aws",
		Description: "Manages an AWS target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetAws,
		},
		Columns: akeylessAkeylessTargetAwsColumns(),
	}
}

func akeylessAkeylessTargetAwsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_key",
			Type:        proto.ColumnType_STRING,
			Description: "AWS secret access key",
		},
		{
			Name:        "access_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "AWS access key ID",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "generate_external_id",
			Type:        proto.ColumnType_BOOL,
			Description: "A unique auto-generated value used in your AWS account when configuring your AWS IAM role to securely delegate access to Akeyless. Relevant only when using GW cloud ID",
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
		{
			Name:        "region",
			Type:        proto.ColumnType_STRING,
			Description: "AWS region",
		},
		{
			Name:        "role_arn",
			Type:        proto.ColumnType_STRING,
			Description: "AWS IAM role identifier that Gateway will assume in your AWS account, relevant only when using external ID",
		},
		{
			Name:        "session_token",
			Type:        proto.ColumnType_STRING,
			Description: "AWS session token (for temporary credentials)",
		},
		{
			Name:        "use_gw_cloud_identity",
			Type:        proto.ColumnType_BOOL,
			Description: "Use gateway cloud identity for authentication",
		},
	}
}

func listAkeylessAkeylessTargetAws(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
