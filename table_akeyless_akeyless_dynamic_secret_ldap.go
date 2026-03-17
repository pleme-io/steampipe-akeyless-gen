package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretLdap() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_ldap",
		Description: "Manages an LDAP dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretLdap,
		},
		Columns: akeylessAkeylessDynamicSecretLdapColumns(),
	}
}

func akeylessAkeylessDynamicSecretLdapColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "ProviderType",
			Type:        proto.ColumnType_STRING,
			Description: "The ProviderType field.",
		},
		{
			Name:        "bind_dn",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP bind DN",
		},
		{
			Name:        "bind_dn_password",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP bind DN password",
		},
		{
			Name:        "custom_username_template",
			Type:        proto.ColumnType_STRING,
			Description: "Customize how temporary usernames are generated using go template",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "external_username",
			Type:        proto.ColumnType_STRING,
			Description: "External username for fixed mode",
		},
		{
			Name:        "fixed_user_claim_keyname",
			Type:        proto.ColumnType_STRING,
			Description: "For externally provided users, denotes the key-name of IdP claim to extract the username from (relevant only for external-username=true)",
		},
		{
			Name:        "group_dn",
			Type:        proto.ColumnType_STRING,
			Description: "Group DN for dynamic users",
		},
		{
			Name:        "host_provider",
			Type:        proto.ColumnType_STRING,
			Description: "Host provider type [explicit/target], Default Host provider is explicit, Relevant only for Secure Remote Access of ssh cert issuer, ldap rotated secret and ldap dynamic secret",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "ldap_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "CA Certificate File Content",
		},
		{
			Name:        "ldap_url",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP server URL",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic secret name",
		},
		{
			Name:        "password_length",
			Type:        proto.ColumnType_STRING,
			Description: "The length of the password to be generated",
		},
		{
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
		},
		{
			Name:        "secure_access_rd_gateway_server",
			Type:        proto.ColumnType_STRING,
			Description: "RD Gateway server",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the producer",
		},
		{
			Name:        "target",
			Type:        proto.ColumnType_JSON,
			Description: "A list of linked targets to be associated, Relevant only for Secure Remote Access for ssh cert issuer, ldap rotated secret and ldap dynamic secret, To specify multiple targets use argument multiple times",
		},
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name associated with this producer",
		},
		{
			Name:        "token_expiration",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP token expiration in seconds",
		},
		{
			Name:        "user_attribute",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP user attribute",
		},
		{
			Name:        "user_dn",
			Type:        proto.ColumnType_STRING,
			Description: "Base DN for user creation",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretLdap(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
