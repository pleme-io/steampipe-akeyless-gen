package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessKmipEnvironment() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_kmip_environment",
		Description: "Manages a KMIP environment in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessKmipEnvironment,
		},
		Columns: akeylessAkeylessKmipEnvironmentColumns(),
	}
}

func akeylessAkeylessKmipEnvironmentColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "certificate_ttl",
			Type:        proto.ColumnType_INT,
			Description: "Server certificate TTL in days",
		},
		{
			Name:        "hostname",
			Type:        proto.ColumnType_STRING,
			Description: "Hostname",
		},
		{
			Name:        "root",
			Type:        proto.ColumnType_STRING,
			Description: "Root path of KMIP Resources",
		},
	}
}

func listAkeylessAkeylessKmipEnvironment(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
