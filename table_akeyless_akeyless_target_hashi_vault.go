package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetHashiVault() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_hashi_vault",
		Description: "Manages a HashiCorp Vault target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetHashiVault,
		},
		Columns: akeylessAkeylessTargetHashiVaultColumns(),
	}
}

func akeylessAkeylessTargetHashiVaultColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "hashi_url",
			Type:        proto.ColumnType_STRING,
			Description: "HashiCorp Vault API URL, e.g. https://vault-mgr01:8200",
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
			Name:        "namespace",
			Type:        proto.ColumnType_JSON,
			Description: "Comma-separated list of vault namespaces",
		},
		{
			Name:        "vault_token",
			Type:        proto.ColumnType_STRING,
			Description: "HashiCorp Vault token",
		},
	}
}

func listAkeylessAkeylessTargetHashiVault(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
