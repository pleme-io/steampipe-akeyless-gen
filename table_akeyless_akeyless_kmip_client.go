package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessKmipClient() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_kmip_client",
		Description: "Manages a KMIP client in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessKmipClient,
		},
		Columns: akeylessAkeylessKmipClientColumns(),
	}
}

func akeylessAkeylessKmipClientColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "activate_keys_on_creation",
			Type:        proto.ColumnType_STRING,
			Description: "If set to 'true', newly created keys on the client will be set to an 'active' state",
		},
		{
			Name:        "certificate_ttl",
			Type:        proto.ColumnType_INT,
			Description: "Client certificate TTL in days",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Client name",
		},
	}
}

func listAkeylessAkeylessKmipClient(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
