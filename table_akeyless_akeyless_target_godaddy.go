package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGodaddy() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_godaddy",
		Description: "Manages a GoDaddy target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGodaddy,
		},
		Columns: akeylessAkeylessTargetGodaddyColumns(),
	}
}

func akeylessAkeylessTargetGodaddyColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "GoDaddy API key",
		},
		{
			Name:        "customer_id",
			Type:        proto.ColumnType_STRING,
			Description: "Customer ID (ShopperId) required for renewal of imported certificates",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "imap_fqdn",
			Type:        proto.ColumnType_STRING,
			Description: "IMAP server FQDN for domain validation",
		},
		{
			Name:        "imap_password",
			Type:        proto.ColumnType_STRING,
			Description: "IMAP password",
		},
		{
			Name:        "imap_port",
			Type:        proto.ColumnType_STRING,
			Description: "IMAP server port",
		},
		{
			Name:        "imap_username",
			Type:        proto.ColumnType_STRING,
			Description: "IMAP username",
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
			Name:        "secret",
			Type:        proto.ColumnType_STRING,
			Description: "GoDaddy API secret",
		},
		{
			Name:        "timeout",
			Type:        proto.ColumnType_STRING,
			Description: "GoDaddy API timeout",
		},
	}
}

func listAkeylessAkeylessTargetGodaddy(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
