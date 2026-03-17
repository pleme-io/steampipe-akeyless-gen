package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetLetsEncrypt() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_lets_encrypt",
		Description: "Manages a Let's Encrypt target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetLetsEncrypt,
		},
		Columns: akeylessAkeylessTargetLetsEncryptColumns(),
	}
}

func akeylessAkeylessTargetLetsEncryptColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "acme_challenge",
			Type:        proto.ColumnType_STRING,
			Description: "The acme_challenge field.",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "dns_target_creds",
			Type:        proto.ColumnType_STRING,
			Description: "Name of existing cloud target for DNS credentials. Required when acme-challenge=dns. Supported: AWS, Azure, GCP targets",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "Email address for ACME account registration",
		},
		{
			Name:        "gcp_project",
			Type:        proto.ColumnType_STRING,
			Description: "GCP Cloud DNS: Project ID. Optional - can be derived from service account",
		},
		{
			Name:        "hosted_zone",
			Type:        proto.ColumnType_STRING,
			Description: "AWS Route53 hosted zone ID. Required when dns-target-creds points to AWS target",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "lets_encrypt_url",
			Type:        proto.ColumnType_STRING,
			Description: "The lets_encrypt_url field.",
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
			Name:        "resource_group",
			Type:        proto.ColumnType_STRING,
			Description: "Azure resource group name. Required when dns-target-creds points to Azure target",
		},
		{
			Name:        "timeout",
			Type:        proto.ColumnType_STRING,
			Description: "The timeout field.",
		},
	}
}

func listAkeylessAkeylessTargetLetsEncrypt(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
