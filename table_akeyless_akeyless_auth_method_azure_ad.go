package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodAzureAd() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_azure_ad",
		Description: "Manages an Azure AD authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodAzureAd,
		},
		Columns: akeylessAkeylessAuthMethodAzureAdColumns(),
	}
}

func akeylessAkeylessAuthMethodAzureAdColumns() []*plugin.Column {
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
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "Azure AD audience (application ID)",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bound_group_id",
			Type:        proto.ColumnType_JSON,
			Description: "Azure AD group IDs to restrict access",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
		},
		{
			Name:        "bound_providers",
			Type:        proto.ColumnType_JSON,
			Description: "Azure resource providers to restrict access",
		},
		{
			Name:        "bound_resource_id",
			Type:        proto.ColumnType_JSON,
			Description: "Azure resource IDs to restrict access",
		},
		{
			Name:        "bound_resource_names",
			Type:        proto.ColumnType_JSON,
			Description: "Azure resource names to restrict access",
		},
		{
			Name:        "bound_resource_types",
			Type:        proto.ColumnType_JSON,
			Description: "Azure resource types to restrict access",
		},
		{
			Name:        "bound_rg_id",
			Type:        proto.ColumnType_JSON,
			Description: "Azure resource group IDs to restrict access",
		},
		{
			Name:        "bound_spid",
			Type:        proto.ColumnType_JSON,
			Description: "Service principal IDs to restrict access",
		},
		{
			Name:        "bound_sub_id",
			Type:        proto.ColumnType_JSON,
			Description: "Azure subscription IDs to restrict access",
		},
		{
			Name:        "bound_tenant_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure tenant ID to restrict access",
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
			Description: "Azure AD issuer URL",
		},
		{
			Name:        "jwks_uri",
			Type:        proto.ColumnType_STRING,
			Description: "JWKS URI for token verification",
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
			Name:        "unique_identifier",
			Type:        proto.ColumnType_STRING,
			Description: "A unique identifier (ID) value which is a \"sub claim\" name that contains details uniquely identifying that resource. This \"sub claim\" is used to distinguish between different identities.",
		},
	}
}

func listAkeylessAkeylessAuthMethodAzureAd(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
