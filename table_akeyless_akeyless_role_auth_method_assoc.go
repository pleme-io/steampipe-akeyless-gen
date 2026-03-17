package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRoleAuthMethodAssoc() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_role_auth_method_assoc",
		Description: "Manages a role-to-auth-method association in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRoleAuthMethodAssoc,
		},
		Columns: akeylessAkeylessRoleAuthMethodAssocColumns(),
	}
}

func akeylessAkeylessRoleAuthMethodAssocColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "am_name",
			Type:        proto.ColumnType_STRING,
			Description: "Auth method name to associate with the role",
		},
		{
			Name:        "case_sensitive",
			Type:        proto.ColumnType_STRING,
			Description: "Case-sensitive sub-claims matching",
		},
		{
			Name:        "role_name",
			Type:        proto.ColumnType_STRING,
			Description: "Role name to associate",
		},
		{
			Name:        "sub_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Sub-claims for the association (key=value pairs)",
		},
	}
}

func listAkeylessAkeylessRoleAuthMethodAssoc(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
