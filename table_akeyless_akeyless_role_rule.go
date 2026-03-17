package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessRoleRule() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_role_rule",
		Description: "Manages a role rule in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessRoleRule,
		},
		Columns: akeylessAkeylessRoleRuleColumns(),
	}
}

func akeylessAkeylessRoleRuleColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "capability",
			Type:        proto.ColumnType_JSON,
			Description: "Permission capability: read, create, update, delete, list, deny",
		},
		{
			Name:        "path",
			Type:        proto.ColumnType_STRING,
			Description: "Item path the rule applies to",
		},
		{
			Name:        "role_name",
			Type:        proto.ColumnType_STRING,
			Description: "Role name to attach the rule to",
		},
		{
			Name:        "rule_type",
			Type:        proto.ColumnType_STRING,
			Description: "Rule type: item-rule or auth-method-rule or role-rule",
		},
		{
			Name:        "ttl",
			Type:        proto.ColumnType_INT,
			Description: "RoleRule ttl",
		},
	}
}

func listAkeylessAkeylessRoleRule(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
