package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetLinked() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_linked",
		Description: "Manages a linked target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetLinked,
		},
		Columns: akeylessAkeylessTargetLinkedColumns(),
	}
}

func akeylessAkeylessTargetLinkedColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "hosts",
			Type:        proto.ColumnType_STRING,
			Description: "Hosts for the linked target",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "parent_target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Parent target name to link to",
		},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "Specifies the hosts type, relevant only when working without parent target",
		},
	}
}

func listAkeylessAkeylessTargetLinked(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
