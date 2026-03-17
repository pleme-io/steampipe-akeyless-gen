package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetSsh() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_ssh",
		Description: "Manages an SSH target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetSsh,
		},
		Columns: akeylessAkeylessTargetSshColumns(),
	}
}

func akeylessAkeylessTargetSshColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "host",
			Type:        proto.ColumnType_STRING,
			Description: "SSH host address",
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
			Name:        "port",
			Type:        proto.ColumnType_STRING,
			Description: "SSH port (default: 22)",
		},
		{
			Name:        "private_key",
			Type:        proto.ColumnType_STRING,
			Description: "SSH private key (PEM)",
		},
		{
			Name:        "private_key_password",
			Type:        proto.ColumnType_STRING,
			Description: "SSH private key passphrase",
		},
		{
			Name:        "ssh_password",
			Type:        proto.ColumnType_STRING,
			Description: "SSH password",
		},
		{
			Name:        "ssh_username",
			Type:        proto.ColumnType_STRING,
			Description: "SSH username",
		},
	}
}

func listAkeylessAkeylessTargetSsh(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
