package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodKerberos() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_kerberos",
		Description: "Manages a Kerberos authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodKerberos,
		},
		Columns: akeylessAkeylessAuthMethodKerberosColumns(),
	}
}

func akeylessAkeylessAuthMethodKerberosColumns() []*plugin.Column {
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
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bind_dn",
			Type:        proto.ColumnType_STRING,
			Description: "The bind_dn field.",
		},
		{
			Name:        "bind_dn_password",
			Type:        proto.ColumnType_STRING,
			Description: "The bind_dn_password field.",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
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
			Name:        "group_attr",
			Type:        proto.ColumnType_STRING,
			Description: "The group_attr field.",
		},
		{
			Name:        "group_dn",
			Type:        proto.ColumnType_STRING,
			Description: "The group_dn field.",
		},
		{
			Name:        "group_filter",
			Type:        proto.ColumnType_STRING,
			Description: "The group_filter field.",
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
			Name:        "keytab_file_data",
			Type:        proto.ColumnType_STRING,
			Description: "Kerberos keytab file data (base64)",
		},
		{
			Name:        "keytab_file_path",
			Type:        proto.ColumnType_STRING,
			Description: "Path to Kerberos keytab file on gateway",
		},
		{
			Name:        "krb5_conf_data",
			Type:        proto.ColumnType_STRING,
			Description: "Kerberos krb5.conf file data (base64)",
		},
		{
			Name:        "krb5_conf_path",
			Type:        proto.ColumnType_STRING,
			Description: "Path to krb5.conf file on gateway",
		},
		{
			Name:        "ldap_anonymous_search",
			Type:        proto.ColumnType_BOOL,
			Description: "The ldap_anonymous_search field.",
		},
		{
			Name:        "ldap_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "The ldap_ca_cert field.",
		},
		{
			Name:        "ldap_url",
			Type:        proto.ColumnType_STRING,
			Description: "The ldap_url field.",
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
			Description: "A unique identifier (ID) value which is a \"sub claim\" name that contains details uniquely identifying that resource. This \"sub claim\" is used to distinguish between different identities.",
		},
		{
			Name:        "user_attribute",
			Type:        proto.ColumnType_STRING,
			Description: "The user_attribute field.",
		},
		{
			Name:        "user_dn",
			Type:        proto.ColumnType_STRING,
			Description: "The user_dn field.",
		},
	}
}

func listAkeylessAkeylessAuthMethodKerberos(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
