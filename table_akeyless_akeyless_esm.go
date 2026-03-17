package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessEsm() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_esm",
		Description: "Manages an external secrets manager (ESM) in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessEsm,
		},
		Columns: akeylessAkeylessEsmColumns(),
	}
}

func akeylessAkeylessEsmColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "binary_value",
			Type:        proto.ColumnType_BOOL,
			Description: "Use this option if the external secret value is a base64 encoded binary",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the external secret",
		},
		{
			Name:        "esm_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the External Secrets Manager item",
		},
		{
			Name:        "secret_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name for the new external secret",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the external secret",
		},
		{
			Name:        "value",
			Type:        proto.ColumnType_STRING,
			Description: "Value of the external secret item, either text or base64 encoded binary",
		},
	}
}

func listAkeylessAkeylessEsm(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
