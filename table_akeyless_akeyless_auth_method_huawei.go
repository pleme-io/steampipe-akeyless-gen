package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodHuawei() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_huawei",
		Description: "Manages a Huawei authentication method in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodHuawei,
		},
		Columns: akeylessAkeylessAuthMethodHuaweiColumns(),
	}
}

func akeylessAkeylessAuthMethodHuaweiColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_expires",
			Type:        proto.ColumnType_INT,
			Description: "Access expiration date in Unix timestamp (select 0 for access without\nexpiry date)",
		},
		{
			Name:        "allowed_client_type",
			Type:        proto.ColumnType_JSON,
			Description: "limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension]",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "auth_url",
			Type:        proto.ColumnType_STRING,
			Description: "sts URL",
		},
		{
			Name:        "bound_domain_id",
			Type:        proto.ColumnType_JSON,
			Description: "A list of domain IDs that the access is restricted to",
		},
		{
			Name:        "bound_domain_name",
			Type:        proto.ColumnType_JSON,
			Description: "A list of domain names that the access is restricted to",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "A CIDR whitelist with the IPs that the access is restricted to",
		},
		{
			Name:        "bound_tenant_id",
			Type:        proto.ColumnType_JSON,
			Description: "A list of full tenant ids that the access is restricted to",
		},
		{
			Name:        "bound_tenant_name",
			Type:        proto.ColumnType_JSON,
			Description: "A list of full tenant names that the access is restricted to",
		},
		{
			Name:        "bound_user_id",
			Type:        proto.ColumnType_JSON,
			Description: "A list of full user ids that the access is restricted to",
		},
		{
			Name:        "bound_user_name",
			Type:        proto.ColumnType_JSON,
			Description: "A list of full user-name that the access is restricted to",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_STRING,
			Description: "Protection from accidental deletion of this object [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Auth Method description",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the auth method would you like to be notified.",
		},
		{
			Name:        "force_sub_claims",
			Type:        proto.ColumnType_BOOL,
			Description: "if true: enforce role-association must include sub claims",
		},
		{
			Name:        "gw_bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "A CIDR whitelist with the GW IPs that the access is restricted to",
		},
		{
			Name:        "jwt_ttl",
			Type:        proto.ColumnType_INT,
			Description: "Jwt TTL",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Auth Method name",
		},
		{
			Name:        "product_type",
			Type:        proto.ColumnType_JSON,
			Description: "Choose the relevant product type for the auth method [sm, sra, pm, dp, ca]",
		},
	}
}

func listAkeylessAkeylessAuthMethodHuawei(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
