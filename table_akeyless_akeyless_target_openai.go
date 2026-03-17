package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetOpenai() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_openai",
		Description: "Manages an OpenAI target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetOpenai,
		},
		Columns: akeylessAkeylessTargetOpenaiColumns(),
	}
}

func akeylessAkeylessTargetOpenaiColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "API key for OpenAI",
		},
		{
			Name:        "api_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "API key ID",
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
			Name:        "model",
			Type:        proto.ColumnType_STRING,
			Description: "Default model to use with OpenAI",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "openai_url",
			Type:        proto.ColumnType_STRING,
			Description: "Base URL of the OpenAI API",
		},
		{
			Name:        "organization_id",
			Type:        proto.ColumnType_STRING,
			Description: "Organization ID",
		},
	}
}

func listAkeylessAkeylessTargetOpenai(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
