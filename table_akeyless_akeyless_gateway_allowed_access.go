package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessGatewayAllowedAccess() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_gateway_allowed_access",
		Description: "Manages gateway allowed access in Akeyless",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessGatewayAllowedAccess,
		},
		Columns: akeylessAkeylessGatewayAllowedAccessColumns(),
	}
}

func akeylessAkeylessGatewayAllowedAccessColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "SubClaimsCaseInsensitive",
			Type:        proto.ColumnType_BOOL,
			Description: "The SubClaimsCaseInsensitive field.",
		},
		{
			Name:        "access_id",
			Type:        proto.ColumnType_STRING,
			Description: "The access ID to attach to this allowed access",
		},
		{
			Name:        "case_sensitive",
			Type:        proto.ColumnType_STRING,
			Description: "Treat sub claims as case-sensitive [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Allowed access description",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Allowed access name",
		},
		{
			Name:        "permissions",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of permissions: defaults, targets, classic_keys, automatic_migration, ldap_auth, dynamic_secret, k8s_auth, log_forwarding, zero_knowledge_encryption, rotated_secret, caching, event_forwarding, admin, kmip, general, rotate_secret_value",
		},
		{
			Name:        "sub_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Key/val of sub claims, e.g. group=admins,developers",
		},
	}
}

func listAkeylessAkeylessGatewayAllowedAccess(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
