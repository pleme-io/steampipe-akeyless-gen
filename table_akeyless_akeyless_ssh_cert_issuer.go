package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessSshCertIssuer() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_ssh_cert_issuer",
		Description: "Manages an SSH certificate issuer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessSshCertIssuer,
		},
		Columns: akeylessAkeylessSshCertIssuerColumns(),
	}
}

func akeylessAkeylessSshCertIssuerColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "ProviderType",
			Type:        proto.ColumnType_STRING,
			Description: "The ProviderType field.",
		},
		{
			Name:        "allowed_users",
			Type:        proto.ColumnType_STRING,
			Description: "Users allowed to fetch the certificate, e.g. root,ubuntu",
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
			Name:        "extensions",
			Type:        proto.ColumnType_JSON,
			Description: "Signed certificates with extensions, e.g. permit-port-forwarding",
		},
		{
			Name:        "external_username",
			Type:        proto.ColumnType_STRING,
			Description: "Externally provided username [true/false]",
		},
		{
			Name:        "fixed_user_claim_keyname",
			Type:        proto.ColumnType_STRING,
			Description: "Key-name of IdP claim to extract username from (for external-username=true)",
		},
		{
			Name:        "host_provider",
			Type:        proto.ColumnType_STRING,
			Description: "Host provider type [explicit/target]",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "SSH certificate issuer name",
		},
		{
			Name:        "principals",
			Type:        proto.ColumnType_STRING,
			Description: "Signed certificates with principal, e.g. example_role1,example_role2",
		},
		{
			Name:        "secure_access_use_internal_ssh_access",
			Type:        proto.ColumnType_BOOL,
			Description: "Use internal SSH Access",
		},
		{
			Name:        "signer_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Key to sign the certificate with",
		},
		{
			Name:        "tag",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the SSH cert issuer",
		},
		{
			Name:        "target",
			Type:        proto.ColumnType_JSON,
			Description: "A list of linked targets to be associated, Relevant only for Secure Remote Access for ssh cert issuer, ldap rotated secret and ldap dynamic secret, To specify multiple targets use argument multiple times",
		},
		{
			Name:        "ttl",
			Type:        proto.ColumnType_INT,
			Description: "Certificate TTL in seconds",
		},
	}
}

func listAkeylessAkeylessSshCertIssuer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
