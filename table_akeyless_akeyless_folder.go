package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessFolder() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_folder",
		Description: "Manages a folder in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessFolder,
		},
		Columns: akeylessAkeylessFolderColumns(),
	}
}

func akeylessAkeylessFolderColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "accessibility",
			Type:        proto.ColumnType_STRING,
			Description: "Folder accessibility: regular or personal",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the folder",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Folder name",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the folder",
		},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "The type field.",
		},
	}
}

func listAkeylessAkeylessFolder(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
