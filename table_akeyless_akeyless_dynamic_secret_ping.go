package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretPing() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_ping",
		Description: "Manages a Ping Identity dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretPing,
		},
		Columns: akeylessAkeylessDynamicSecretPingColumns(),
	}
}

func akeylessAkeylessDynamicSecretPingColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic secret name",
		},
		{
			Name:        "ping_administrative_port",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity administrative port",
		},
		{
			Name:        "ping_atm_id",
			Type:        proto.ColumnType_STRING,
			Description: "Set a specific Access Token Management (ATM) instance for the created OAuth Client by providing the ATM Id. If no explicit value is given, the default pingfederate server ATM will be set.",
		},
		{
			Name:        "ping_authorization_port",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Federate authorization port",
		},
		{
			Name:        "ping_cert_subject_dn",
			Type:        proto.ColumnType_STRING,
			Description: "The subject DN of the client certificate. If no explicit value is given, the producer will create CA certificate and matched client certificate and return it as value. Used in conjunction with ping-issuer-dn (relevant for CLIENT_TLS_CERTIFICATE authentication method)",
		},
		{
			Name:        "ping_client_authentication_type",
			Type:        proto.ColumnType_STRING,
			Description: "OAuth Client Authentication Type [CLIENT_SECRET, PRIVATE_KEY_JWT, CLIENT_TLS_CERTIFICATE]",
		},
		{
			Name:        "ping_enforce_replay_prevention",
			Type:        proto.ColumnType_STRING,
			Description: "Determines whether PingFederate requires a unique signed JWT from the client for each action (relevant for PRIVATE_KEY_JWT authentication method) [true/false]",
		},
		{
			Name:        "ping_grant_types",
			Type:        proto.ColumnType_JSON,
			Description: "Ping Identity grant types",
		},
		{
			Name:        "ping_issuer_dn",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity issuer DN",
		},
		{
			Name:        "ping_jwks",
			Type:        proto.ColumnType_STRING,
			Description: "Base64-encoded JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT (Sign Algo: RS256) and return it as value (relevant for PRIVATE_KEY_JWT authentication method)",
		},
		{
			Name:        "ping_jwks_url",
			Type:        proto.ColumnType_STRING,
			Description: "The URL of the JSON Web Key Set (JWKS). If no explicit value is given, the producer will create JWKs and matched signed JWT and return it as value (relevant for PRIVATE_KEY_JWT authentication method)",
		},
		{
			Name:        "ping_password",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity password",
		},
		{
			Name:        "ping_privileged_user",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity privileged username",
		},
		{
			Name:        "ping_redirect_uris",
			Type:        proto.ColumnType_JSON,
			Description: "Ping Identity redirect URIs",
		},
		{
			Name:        "ping_restricted_scopes",
			Type:        proto.ColumnType_JSON,
			Description: "Limit the OAuth client to specific scopes list",
		},
		{
			Name:        "ping_signing_algo",
			Type:        proto.ColumnType_STRING,
			Description: "The signing algorithm that the client must use to sign its request objects [RS256,RS384,RS512,ES256,ES384,ES512,PS256,PS384,PS512] If no explicit value is given, the client can use any of the supported signing algorithms (relevant for PRIVATE_KEY_JWT authentication method)",
		},
		{
			Name:        "ping_url",
			Type:        proto.ColumnType_STRING,
			Description: "Ping Identity server URL",
		},
		{
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "tags",
			Type:        proto.ColumnType_JSON,
			Description: "Tags for the producer",
		},
		{
			Name:        "target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name associated with this producer",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretPing(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
