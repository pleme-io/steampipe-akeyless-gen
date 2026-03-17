package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_group",
		Description: "Manages a group in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessGroup,
		},
		Columns: akeylessAkeylessGroupColumns(),
	}
}

func akeylessAkeylessGroupColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the group",
		},
		{
			Name:        "group_alias",
			Type:        proto.ColumnType_STRING,
			Description: "A short group alias",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Group name",
		},
		{
			Name:        "user_assignment",
			Type:        proto.ColumnType_STRING,
			Description: "JSON string defining access permission assignment",
		},
	}
}

func listAkeylessAkeylessGroup(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
