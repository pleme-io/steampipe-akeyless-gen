package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessClassicKey() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_classic_key",
		Description: "Manages a classic cryptographic key",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessClassicKey,
		},
		Columns: akeylessAkeylessClassicKeyColumns(),
	}
}

func akeylessAkeylessClassicKeyColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "alg",
			Type:        proto.ColumnType_STRING,
			Description: "Classic key algorithm: AES128GCM, AES256GCM, AES128SIV, AES256SIV, RSA1024, RSA2048, RSA3072, RSA4096, EC256, EC384",
		},
		{
			Name:        "auto_rotate",
			Type:        proto.ColumnType_STRING,
			Description: "Whether to automatically rotate every rotation_interval days, or disable existing automatic rotation [true/false]",
		},
		{
			Name:        "cert_file_data",
			Type:        proto.ColumnType_STRING,
			Description: "Certificate PEM data (for RSA keys)",
		},
		{
			Name:        "certificate_common_name",
			Type:        proto.ColumnType_STRING,
			Description: "Common name for generated certificate",
		},
		{
			Name:        "certificate_country",
			Type:        proto.ColumnType_STRING,
			Description: "Country for generated certificate",
		},
		{
			Name:        "certificate_digest_algo",
			Type:        proto.ColumnType_STRING,
			Description: "Digest algorithm to be used for the certificate key signing.",
		},
		{
			Name:        "certificate_format",
			Type:        proto.ColumnType_STRING,
			Description: "The certificate_format field.",
		},
		{
			Name:        "certificate_locality",
			Type:        proto.ColumnType_STRING,
			Description: "Locality for generated certificate",
		},
		{
			Name:        "certificate_organization",
			Type:        proto.ColumnType_STRING,
			Description: "Organization for generated certificate",
		},
		{
			Name:        "certificate_province",
			Type:        proto.ColumnType_STRING,
			Description: "Province for generated certificate",
		},
		{
			Name:        "certificate_ttl",
			Type:        proto.ColumnType_INT,
			Description: "Certificate TTL in days",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Key description",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the certificate would you like to be notified.",
		},
		{
			Name:        "generate_self_signed_certificate",
			Type:        proto.ColumnType_BOOL,
			Description: "Generate a self-signed certificate",
		},
		{
			Name:        "gpg_alg",
			Type:        proto.ColumnType_STRING,
			Description: "gpg alg: Relevant only if GPG key type selected; options: [RSA1024, RSA2048, RSA3072, RSA4096, Ed25519]",
		},
		{
			Name:        "hash_algorithm",
			Type:        proto.ColumnType_STRING,
			Description: "Specifies the hash algorithm used for the encryption key's operations, available options: [SHA256, SHA384, SHA512]",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "key_data",
			Type:        proto.ColumnType_STRING,
			Description: "Key material (base64 encoded)",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "ClassicKey name",
		},
		{
			Name:        "protection_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Encryption key used to protect this key",
		},
		{
			Name:        "rotation_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the rotation of the item would you like to be notified",
		},
		{
			Name:        "rotation_interval",
			Type:        proto.ColumnType_STRING,
			Description: "The number of days to wait between every automatic rotation (1-365)",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the key",
		},
	}
}

func listAkeylessAkeylessClassicKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
