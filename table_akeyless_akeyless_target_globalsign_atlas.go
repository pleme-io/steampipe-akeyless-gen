package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGlobalsignAtlas() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_globalsign_atlas",
		Description: "Manages a GlobalSign Atlas target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGlobalsignAtlas,
		},
		Columns: akeylessAkeylessTargetGlobalsignAtlasColumns(),
	}
}

func akeylessAkeylessTargetGlobalsignAtlasColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "GlobalSign Atlas API key",
		},
		{
			Name:        "api_secret",
			Type:        proto.ColumnType_STRING,
			Description: "GlobalSign Atlas API secret",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
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
			Name:        "mtls_cert_data_base64",
			Type:        proto.ColumnType_STRING,
			Description: "Mutual TLS Certificate contents of the GlobalSign Atlas account encoded in base64, either mtls-cert-file-path or mtls-cert-data-base64 must be supplied",
		},
		{
			Name:        "mtls_key_data_base64",
			Type:        proto.ColumnType_STRING,
			Description: "Mutual TLS Key contents of the GlobalSign Atlas account encoded in base64, either mtls-key-file-path or mtls-data-base64 must be supplied",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "timeout",
			Type:        proto.ColumnType_STRING,
			Description: "Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.",
		},
	}
}

func listAkeylessAkeylessTargetGlobalsignAtlas(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
