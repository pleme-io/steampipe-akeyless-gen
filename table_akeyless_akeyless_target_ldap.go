package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetLdap() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_ldap",
		Description: "Manages an LDAP target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetLdap,
		},
		Columns: akeylessAkeylessTargetLdapColumns(),
	}
}

func akeylessAkeylessTargetLdapColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "ldap_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "CA Certificate File Content",
		},
		{
			Name:        "ldap_url",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP server URL (e.g., ldap://host:389)",
		},
		{
			Name:        "max_versions",
			Type:        proto.ColumnType_STRING,
			Description: "Set the maximum number of versions, limited by the account settings defaults.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "server_type",
			Type:        proto.ColumnType_STRING,
			Description: "Set Ldap server type, Options:[OpenLDAP, ActiveDirectory]. Default is OpenLDAP",
		},
		{
			Name:        "token_expiration",
			Type:        proto.ColumnType_STRING,
			Description: "LDAP token expiration in seconds",
		},
	}
}

func listAkeylessAkeylessTargetLdap(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
