package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodCert() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_cert",
		Description: "Manages a certificate authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodCert,
		},
		Columns: akeylessAkeylessAuthMethodCertColumns(),
	}
}

func akeylessAkeylessAuthMethodCertColumns() []*plugin.Column {
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
			Name:        "allowed_cors",
			Type:        proto.ColumnType_STRING,
			Description: "Comma separated list of allowed CORS domains to be validated as part of the authentication flow.",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bound_common_names",
			Type:        proto.ColumnType_JSON,
			Description: "Bound common names for certificate matching",
		},
		{
			Name:        "bound_dns_sans",
			Type:        proto.ColumnType_JSON,
			Description: "Bound DNS SANs for certificate matching",
		},
		{
			Name:        "bound_email_sans",
			Type:        proto.ColumnType_JSON,
			Description: "Bound email SANs for certificate matching",
		},
		{
			Name:        "bound_extensions",
			Type:        proto.ColumnType_JSON,
			Description: "Bound extensions (OID:value pairs)",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
		},
		{
			Name:        "bound_organizational_units",
			Type:        proto.ColumnType_JSON,
			Description: "Bound organizational units for certificate matching",
		},
		{
			Name:        "bound_uri_sans",
			Type:        proto.ColumnType_JSON,
			Description: "Bound URI SANs for certificate matching",
		},
		{
			Name:        "certificate_data",
			Type:        proto.ColumnType_STRING,
			Description: "PEM certificate data for validation",
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
			Name:        "require_crl_dp",
			Type:        proto.ColumnType_BOOL,
			Description: "Require certificate CRL distribution points (CDP) and enforce CRL validation during authentication.",
		},
		{
			Name:        "revoked_cert_ids",
			Type:        proto.ColumnType_JSON,
			Description: "Revoked certificate serial numbers",
		},
		{
			Name:        "unique_identifier",
			Type:        proto.ColumnType_STRING,
			Description: "A unique identifier (ID) value should be configured, such as common_name or organizational_unit\nWhenever a user logs in with a token, these authentication types issue\na \"sub claim\" that contains details uniquely identifying that user.\nThis sub claim includes a key containing the ID value that you\nconfigured, and is used to distinguish between different users from\nwithin the same organization.",
		},
	}
}

func listAkeylessAkeylessAuthMethodCert(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
