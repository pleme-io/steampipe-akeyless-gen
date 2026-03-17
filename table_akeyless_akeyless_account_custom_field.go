package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAccountCustomField() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_account_custom_field",
		Description: "Manages an account custom field in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAccountCustomField,
		},
		Columns: akeylessAkeylessAccountCustomFieldColumns(),
	}
}

func akeylessAkeylessAccountCustomFieldColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the custom field",
		},
		{
			Name:        "object",
			Type:        proto.ColumnType_STRING,
			Description: "The object to create the custom field",
		},
		{
			Name:        "object_type",
			Type:        proto.ColumnType_STRING,
			Description: "The object type to create the custom field [e.g. STATIC_SECRET,DYNAMIC_SECRET,ROTATED_SECRET]",
		},
		{
			Name:        "required",
			Type:        proto.ColumnType_BOOL,
			Description: "Specify whether the custom field is mandatory",
		},
	}
}

func listAkeylessAkeylessAccountCustomField(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
