package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetWindows() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_windows",
		Description: "Manages a Windows target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetWindows,
		},
		Columns: akeylessAkeylessTargetWindowsColumns(),
	}
}

func akeylessAkeylessTargetWindowsColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "certificate",
			Type:        proto.ColumnType_STRING,
			Description: "Client certificate (PEM)",
		},
		{
			Name:        "connection_type",
			Type:        proto.ColumnType_STRING,
			Description: "Type of connection to Windows Server [credentials/parent-target]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "domain",
			Type:        proto.ColumnType_STRING,
			Description: "Windows domain",
		},
		{
			Name:        "hostname",
			Type:        proto.ColumnType_STRING,
			Description: "Windows host address",
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
			Name:        "parent_target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the parent target, relevant only when connection-type is parent-target",
		},
		{
			Name:        "port",
			Type:        proto.ColumnType_STRING,
			Description: "WinRM port (default: 5986)",
		},
		{
			Name:        "use_tls",
			Type:        proto.ColumnType_BOOL,
			Description: "Use TLS for WinRM connection",
		},
	}
}

func listAkeylessAkeylessTargetWindows(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
