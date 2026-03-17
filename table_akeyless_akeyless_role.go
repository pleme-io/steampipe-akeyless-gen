package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRole() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_role",
		Description: "Manages a role in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRole,
		},
		Columns: akeylessAkeylessRoleColumns(),
	}
}

func akeylessAkeylessRoleColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "analytics_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view analytics",
		},
		{
			Name:        "audit_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view audit logs",
		},
		{
			Name:        "comment",
			Type:        proto.ColumnType_STRING,
			Description: "Deprecated - use description",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_STRING,
			Description: "Protection from accidental deletion of this object [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Role description",
		},
		{
			Name:        "event_center_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view Event Center. Currently only 'none', 'scoped' and 'all'\nvalues are supported",
		},
		{
			Name:        "event_forwarders_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to manage Event Forwarders. Currently only 'none' and 'all' values are supported.",
		},
		{
			Name:        "event_forwarders_name",
			Type:        proto.ColumnType_JSON,
			Description: "Allow this role to manage the following Event Forwarders.",
		},
		{
			Name:        "gw_analytics_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view gateway analytics",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Role name",
		},
		{
			Name:        "reverse_rbac_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view Reverse RBAC. Supported values: 'scoped', 'all'.",
		},
		{
			Name:        "sra_reports_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view SRA reports",
		},
		{
			Name:        "usage_reports_access",
			Type:        proto.ColumnType_STRING,
			Description: "Allow this role to view Usage Report. Currently only 'none' and\n'all' values are supported.",
		},
	}
}

func listAkeylessAkeylessRole(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
