package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodOidc() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_oidc",
		Description: "Manages an OIDC authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodOidc,
		},
		Columns: akeylessAkeylessAuthMethodOidcColumns(),
	}
}

func akeylessAkeylessAuthMethodOidcColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_expires",
			Type:        proto.ColumnType_INT,
			Description: "Access expiration in Unix time (0 = no expiry)",
		},
		{
			Name:        "allowed_client_type",
			Type:        proto.ColumnType_JSON,
			Description: "limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension]",
		},
		{
			Name:        "allowed_redirect_uri",
			Type:        proto.ColumnType_JSON,
			Description: "Allowed redirect URIs after OIDC login",
		},
		{
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "OIDC audience",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
		},
		{
			Name:        "client_id",
			Type:        proto.ColumnType_STRING,
			Description: "OIDC client ID",
		},
		{
			Name:        "client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "OIDC client secret",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
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
			Description: "Force sub-claims enforcement",
		},
		{
			Name:        "gw_bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "Gateway CIDR whitelist",
		},
		{
			Name:        "issuer",
			Type:        proto.ColumnType_STRING,
			Description: "OIDC issuer URL",
		},
		{
			Name:        "jwt_ttl",
			Type:        proto.ColumnType_INT,
			Description: "JWT TTL in seconds",
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
		{
			Name:        "required_scopes",
			Type:        proto.ColumnType_JSON,
			Description: "Required OIDC scopes (comma-separated)",
		},
		{
			Name:        "required_scopes_prefix",
			Type:        proto.ColumnType_STRING,
			Description: "Prefix for required scopes",
		},
		{
			Name:        "subclaims_delimiters",
			Type:        proto.ColumnType_JSON,
			Description: "Delimiters for subclaim splitting",
		},
		{
			Name:        "unique_identifier",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier claim name in the OIDC token",
		},
	}
}

func listAkeylessAkeylessAuthMethodOidc(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
