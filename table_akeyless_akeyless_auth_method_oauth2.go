package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodOauth2() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_oauth2",
		Description: "Manages an OAuth 2.0 authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodOauth2,
		},
		Columns: akeylessAkeylessAuthMethodOauth2Columns(),
	}
}

func akeylessAkeylessAuthMethodOauth2Columns() []*plugin.Column {
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
			Description: "OAuth2 audience",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bound_client_ids",
			Type:        proto.ColumnType_JSON,
			Description: "OAuth2 client IDs to restrict access",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
		},
		{
			Name:        "cert",
			Type:        proto.ColumnType_STRING,
			Description: "CertificateFile Path to a file that contain the certificate in a PEM format.",
		},
		{
			Name:        "cert_file_data",
			Type:        proto.ColumnType_STRING,
			Description: "CertificateFileData PEM Certificate in a Base64 format.",
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
			Name:        "gateway_url",
			Type:        proto.ColumnType_STRING,
			Description: "Akeyless Gateway URL (Configuration Management port). Relevant only when the jwks-uri is accessible only from the gateway.",
		},
		{
			Name:        "gw_bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "Gateway CIDR whitelist",
		},
		{
			Name:        "issuer",
			Type:        proto.ColumnType_STRING,
			Description: "OAuth2 issuer URL",
		},
		{
			Name:        "jwks_json_data",
			Type:        proto.ColumnType_STRING,
			Description: "JWKS JSON data for token verification",
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
			Name:        "subclaims_delimiters",
			Type:        proto.ColumnType_JSON,
			Description: "A list of additional sub claims delimiters (relevant only for SAML, OIDC, OAuth2/JWT)",
		},
		{
			Name:        "unique_identifier",
			Type:        proto.ColumnType_STRING,
			Description: "Unique identifier claim name in the JWT token",
		},
	}
}

func listAkeylessAkeylessAuthMethodOauth2(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
