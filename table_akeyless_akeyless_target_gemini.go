package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGemini() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_gemini",
		Description: "Manages a Gemini target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGemini,
		},
		Columns: akeylessAkeylessTargetGeminiColumns(),
	}
}

func akeylessAkeylessTargetGeminiColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "api_key",
			Type:        proto.ColumnType_STRING,
			Description: "API key for Gemini",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "gemini_url",
			Type:        proto.ColumnType_STRING,
			Description: "Base URL of the Gemini API",
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
	}
}

func listAkeylessAkeylessTargetGemini(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
