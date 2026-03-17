package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetSectigo() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_sectigo",
		Description: "Manages a Sectigo target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetSectigo,
		},
		Columns: akeylessAkeylessTargetSectigoColumns(),
	}
}

func akeylessAkeylessTargetSectigoColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "certificate_profile_id",
			Type:        proto.ColumnType_INT,
			Description: "Certificate Profile ID in Sectigo account",
		},
		{
			Name:        "customer_uri",
			Type:        proto.ColumnType_STRING,
			Description: "Customer Uri of the Sectigo account",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "external_requester",
			Type:        proto.ColumnType_STRING,
			Description: "External Requester - a comma separated list of emails",
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
			Name:        "organization_id",
			Type:        proto.ColumnType_INT,
			Description: "Organization ID in Sectigo account",
		},
		{
			Name:        "timeout",
			Type:        proto.ColumnType_STRING,
			Description: "Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.",
		},
	}
}

func listAkeylessAkeylessTargetSectigo(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
