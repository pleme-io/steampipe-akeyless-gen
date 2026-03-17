package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessPkiCertIssuer() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_pki_cert_issuer",
		Description: "Manages a PKI certificate issuer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessPkiCertIssuer,
		},
		Columns: akeylessAkeylessPkiCertIssuerColumns(),
	}
}

func akeylessAkeylessPkiCertIssuerColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "allow_any_name",
			Type:        proto.ColumnType_BOOL,
			Description: "If set, clients can request certificates for any CN",
		},
		{
			Name:        "allow_copy_ext_from_csr",
			Type:        proto.ColumnType_BOOL,
			Description: "Allow copying extra extensions from CSR",
		},
		{
			Name:        "allow_subdomains",
			Type:        proto.ColumnType_BOOL,
			Description: "Allow certificates for subdomains of allowed domains",
		},
		{
			Name:        "allowed_domains",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-delimited list of allowed domains",
		},
		{
			Name:        "allowed_extra_extensions",
			Type:        proto.ColumnType_STRING,
			Description: "JSON string of allowed extra extensions",
		},
		{
			Name:        "allowed_ip_sans",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-delimited list of allowed CIDRs for IP SANs",
		},
		{
			Name:        "allowed_uri_sans",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-delimited list of allowed URIs for URI SANs",
		},
		{
			Name:        "auto_renew",
			Type:        proto.ColumnType_BOOL,
			Description: "Automatically renew certificates before expiration",
		},
		{
			Name:        "ca_target",
			Type:        proto.ColumnType_STRING,
			Description: "Name of existing CA target (required in Public CA mode)",
		},
		{
			Name:        "client_flag",
			Type:        proto.ColumnType_BOOL,
			Description: "Flag certificates for client auth use",
		},
		{
			Name:        "code_signing_flag",
			Type:        proto.ColumnType_BOOL,
			Description: "Flag certificates for code signing use",
		},
		{
			Name:        "country",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of countries for issued certificate",
		},
		{
			Name:        "create_private_crl",
			Type:        proto.ColumnType_BOOL,
			Description: "Expose CRL endpoint in Gateway",
		},
		{
			Name:        "create_private_ocsp",
			Type:        proto.ColumnType_BOOL,
			Description: "Expose private OCSP endpoint",
		},
		{
			Name:        "create_public_crl",
			Type:        proto.ColumnType_BOOL,
			Description: "Expose public CRL endpoint",
		},
		{
			Name:        "create_public_ocsp",
			Type:        proto.ColumnType_BOOL,
			Description: "Expose public OCSP endpoint",
		},
		{
			Name:        "critical_key_usage",
			Type:        proto.ColumnType_STRING,
			Description: "Mark key usage as critical",
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
			Name:        "destination_path",
			Type:        proto.ColumnType_STRING,
			Description: "Destination path for generated certificates",
		},
		{
			Name:        "disable_wildcards",
			Type:        proto.ColumnType_BOOL,
			Description: "Disable wildcard certificates",
		},
		{
			Name:        "enable_acme",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable ACME protocol",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the certificate would you like to be notified.",
		},
		{
			Name:        "gw_cluster_url",
			Type:        proto.ColumnType_STRING,
			Description: "Gateway cluster URL",
		},
		{
			Name:        "is_ca",
			Type:        proto.ColumnType_BOOL,
			Description: "If set, the basic constraints is-ca flag is set",
		},
		{
			Name:        "key_usage",
			Type:        proto.ColumnType_STRING,
			Description: "Key usage: DigitalSignature, KeyAgreement, KeyEncipherment",
		},
		{
			Name:        "locality",
			Type:        proto.ColumnType_STRING,
			Description: "Locality for issued certificate",
		},
		{
			Name:        "max_path_len",
			Type:        proto.ColumnType_INT,
			Description: "Maximum path length for CA certificates",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "PKI certificate issuer name",
		},
		{
			Name:        "not_enforce_hostnames",
			Type:        proto.ColumnType_BOOL,
			Description: "Do not enforce hostnames",
		},
		{
			Name:        "not_require_cn",
			Type:        proto.ColumnType_BOOL,
			Description: "Do not require common name",
		},
		{
			Name:        "ocsp_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "OCSP response TTL",
		},
		{
			Name:        "organizational_units",
			Type:        proto.ColumnType_STRING,
			Description: "Organizational units for issued certificate",
		},
		{
			Name:        "organizations",
			Type:        proto.ColumnType_STRING,
			Description: "Organization for issued certificate",
		},
		{
			Name:        "postal_code",
			Type:        proto.ColumnType_STRING,
			Description: "Postal code for issued certificate",
		},
		{
			Name:        "protect_certificates",
			Type:        proto.ColumnType_BOOL,
			Description: "Protect generated certificates",
		},
		{
			Name:        "province",
			Type:        proto.ColumnType_STRING,
			Description: "Province for issued certificate",
		},
		{
			Name:        "scheduled_renew",
			Type:        proto.ColumnType_INT,
			Description: "Number of days before expiration to renew certificates",
		},
		{
			Name:        "server_flag",
			Type:        proto.ColumnType_BOOL,
			Description: "Flag certificates for server auth use",
		},
		{
			Name:        "signer_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Key to sign certificates with",
		},
		{
			Name:        "street_address",
			Type:        proto.ColumnType_STRING,
			Description: "Street address for issued certificate",
		},
		{
			Name:        "tag",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the PKI cert issuer",
		},
		{
			Name:        "ttl",
			Type:        proto.ColumnType_STRING,
			Description: "Certificate TTL",
		},
	}
}

func listAkeylessAkeylessPkiCertIssuer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
