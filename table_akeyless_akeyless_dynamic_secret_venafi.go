package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretVenafi() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_venafi",
		Description: "Manages a Venafi dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretVenafi,
		},
		Columns: akeylessAkeylessDynamicSecretVenafiColumns(),
	}
}

func akeylessAkeylessDynamicSecretVenafiColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "admin_rotation_interval_days",
			Type:        proto.ColumnType_INT,
			Description: "Admin credentials rotation interval (days)",
		},
		{
			Name:        "allow_subdomains",
			Type:        proto.ColumnType_BOOL,
			Description: "Allow subdomains",
		},
		{
			Name:        "allowed_domains",
			Type:        proto.ColumnType_JSON,
			Description: "Allowed domains",
		},
		{
			Name:        "auto_generated_folder",
			Type:        proto.ColumnType_STRING,
			Description: "Auto generated folder",
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
			Name:        "enable_admin_rotation",
			Type:        proto.ColumnType_BOOL,
			Description: "Automatic admin credentials rotation",
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
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "root_first_in_chain",
			Type:        proto.ColumnType_BOOL,
			Description: "Root first in chain",
		},
		{
			Name:        "sign_using_akeyless_pki",
			Type:        proto.ColumnType_BOOL,
			Description: "Use Akeyless PKI issuer or Venafi issuer",
		},
		{
			Name:        "signer_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Signer key name",
		},
		{
			Name:        "store_private_key",
			Type:        proto.ColumnType_BOOL,
			Description: "Store private key in Akeyless",
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
		{
			Name:        "venafi_access_token",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi Access Token to use to access the TPP environment (Relevant when using TPP)",
		},
		{
			Name:        "venafi_api_key",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi API key",
		},
		{
			Name:        "venafi_baseurl",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi base URL",
		},
		{
			Name:        "venafi_client_id",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi Client ID that was used when the access token was generated",
		},
		{
			Name:        "venafi_refresh_token",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi Refresh Token to use when the Access Token is expired (Relevant when using TPP)",
		},
		{
			Name:        "venafi_use_tpp",
			Type:        proto.ColumnType_BOOL,
			Description: "Use Venafi TPP (instead of Cloud)",
		},
		{
			Name:        "venafi_zone",
			Type:        proto.ColumnType_STRING,
			Description: "Venafi zone/policy folder",
		},
	}
}

func listAkeylessAkeylessDynamicSecretVenafi(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
