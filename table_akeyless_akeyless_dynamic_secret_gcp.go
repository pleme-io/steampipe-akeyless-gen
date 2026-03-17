package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretGcp() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_gcp",
		Description: "Manages a GCP dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretGcp,
		},
		Columns: akeylessAkeylessDynamicSecretGcpColumns(),
	}
}

func akeylessAkeylessDynamicSecretGcpColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_type",
			Type:        proto.ColumnType_STRING,
			Description: "The access_type field.",
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
			Name:        "fixed_user_claim_keyname",
			Type:        proto.ColumnType_STRING,
			Description: "For externally provided users, denotes the key-name of IdP claim to extract the username from (Relevant only when --access-type=external)",
		},
		{
			Name:        "gcp_cred_type",
			Type:        proto.ColumnType_STRING,
			Description: "GCP credential type: token or key",
		},
		{
			Name:        "gcp_key",
			Type:        proto.ColumnType_STRING,
			Description: "Base64-encoded service account private key text",
		},
		{
			Name:        "gcp_key_algo",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account key algorithm",
		},
		{
			Name:        "gcp_project_id",
			Type:        proto.ColumnType_STRING,
			Description: "GCP Project ID override for dynamic secret operations",
		},
		{
			Name:        "gcp_sa_email",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account email",
		},
		{
			Name:        "gcp_token_scopes",
			Type:        proto.ColumnType_STRING,
			Description: "GCP access token scopes (comma-separated)",
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
			Name:        "role_binding",
			Type:        proto.ColumnType_STRING,
			Description: "Role binding in the format Role,Resource",
		},
		{
			Name:        "role_names",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of GCP roles to assign to the user (Relevant only when --access-type=external)",
		},
		{
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
		},
		{
			Name:        "service_account_type",
			Type:        proto.ColumnType_STRING,
			Description: "Service account type: fixed or dynamic",
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

func listAkeylessAkeylessDynamicSecretGcp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
