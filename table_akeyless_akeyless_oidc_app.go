package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessOidcApp() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_oidc_app",
		Description: "Manages an OIDC application in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessOidcApp,
		},
		Columns: akeylessAkeylessOidcAppColumns(),
	}
}

func akeylessAkeylessOidcAppColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "accessibility",
			Type:        proto.ColumnType_STRING,
			Description: "for personal password manager",
		},
		{
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "A comma separated list of allowed audiences",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_STRING,
			Description: "Protection from accidental deletion of this object [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the OIDC application (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "OIDC application name",
		},
		{
			Name:        "permission_assignment",
			Type:        proto.ColumnType_STRING,
			Description: "A json string defining the access permission assignment for this app",
		},
		{
			Name:        "public",
			Type:        proto.ColumnType_BOOL,
			Description: "Set to true if the app is public (cannot keep secrets)",
		},
		{
			Name:        "redirect_uris",
			Type:        proto.ColumnType_STRING,
			Description: "A comma separated list of allowed redirect uris",
		},
		{
			Name:        "scopes",
			Type:        proto.ColumnType_STRING,
			Description: "A comma separated list of allowed scopes",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Add tags attached to this object",
		},
	}
}

func listAkeylessAkeylessOidcApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
