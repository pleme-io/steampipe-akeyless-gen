package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretGoogleWorkspace() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_google_workspace",
		Description: "Manages a Google Workspace dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretGoogleWorkspace,
		},
		Columns: akeylessAkeylessDynamicSecretGoogleWorkspaceColumns(),
	}
}

func akeylessAkeylessDynamicSecretGoogleWorkspaceColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_mode",
			Type:        proto.ColumnType_STRING,
			Description: "The access_mode field.",
		},
		{
			Name:        "admin_email",
			Type:        proto.ColumnType_STRING,
			Description: "Admin user email",
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
			Description: "For externally provided users, denotes the key-name of IdP claim to extract the username from",
		},
		{
			Name:        "gcp_key",
			Type:        proto.ColumnType_STRING,
			Description: "Base64-encoded service account private key text",
		},
		{
			Name:        "group_email",
			Type:        proto.ColumnType_STRING,
			Description: "A group email, relevant only for group access-mode",
		},
		{
			Name:        "group_role",
			Type:        proto.ColumnType_STRING,
			Description: "The group_role field.",
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
			Name:        "role_name",
			Type:        proto.ColumnType_STRING,
			Description: "Google Workspace admin role",
		},
		{
			Name:        "role_scope",
			Type:        proto.ColumnType_STRING,
			Description: "Google Workspace role scope (ORG_UNIT)",
		},
		{
			Name:        "secure_access_delay",
			Type:        proto.ColumnType_INT,
			Description: "The delay duration, in seconds, to wait after generating just-in-time credentials. Accepted range: 0-120 seconds",
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

func listAkeylessAkeylessDynamicSecretGoogleWorkspace(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
