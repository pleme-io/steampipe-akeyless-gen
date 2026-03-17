package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessCertificate() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_certificate",
		Description: "Manages a certificate in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessCertificate,
		},
		Columns: akeylessAkeylessCertificateColumns(),
	}
}

func akeylessAkeylessCertificateColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "certificate_data",
			Type:        proto.ColumnType_STRING,
			Description: "Content of the certificate in a Base64 format.",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_STRING,
			Description: "Protection from accidental deletion of this object [true/false]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the certificate would you like to be notified.",
		},
		{
			Name:        "format",
			Type:        proto.ColumnType_STRING,
			Description: "CertificateFormat of the certificate and private key, possible values: cer,crt,pem,pfx,p12.\nRequired when passing inline certificate content with --certificate-data or --key-data, otherwise format is derived from the file extension.",
		},
		{
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key to use to encrypt the certificate's key (if empty, the\naccount default protectionKey key will be used)",
		},
		{
			Name:        "key_data",
			Type:        proto.ColumnType_STRING,
			Description: "Content of the certificate's private key in a Base64 format.",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Certificate name",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Add tags attached to this object",
		},
	}
}

func listAkeylessAkeylessCertificate(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
