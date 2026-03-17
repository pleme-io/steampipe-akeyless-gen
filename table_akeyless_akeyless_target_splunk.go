package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetSplunk() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_splunk",
		Description: "Manages a Splunk target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetSplunk,
		},
		Columns: akeylessAkeylessTargetSplunkColumns(),
	}
}

func akeylessAkeylessTargetSplunkColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "Splunk token audience (required when using token authentication for rotation)",
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
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "token_owner",
			Type:        proto.ColumnType_STRING,
			Description: "Splunk Token Owner (required when using token authentication for rotation)",
		},
		{
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Description: "Splunk server URL",
		},
		{
			Name:        "use_tls",
			Type:        proto.ColumnType_BOOL,
			Description: "Use TLS certificate verification when connecting to the Splunk management API",
		},
	}
}

func listAkeylessAkeylessTargetSplunk(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
