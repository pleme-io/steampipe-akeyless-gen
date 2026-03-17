package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretRdp() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_rdp",
		Description: "Manages an RDP dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretRdp,
		},
		Columns: akeylessAkeylessDynamicSecretRdpColumns(),
	}
}

func akeylessAkeylessDynamicSecretRdpColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "allow_user_extend_session",
			Type:        proto.ColumnType_INT,
			Description: "AllowUserExtendSession",
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
			Description: "For externally provided users, denotes the key-name of IdP claim to extract the username from (relevant only for fixed-user-only=true)",
		},
		{
			Name:        "fixed_user_only",
			Type:        proto.ColumnType_STRING,
			Description: "Allow access using externally (IdP) provided username [true/false]",
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
			Name:        "rdp_admin_name",
			Type:        proto.ColumnType_STRING,
			Description: "RDP admin username",
		},
		{
			Name:        "rdp_admin_pwd",
			Type:        proto.ColumnType_STRING,
			Description: "RDP admin password",
		},
		{
			Name:        "rdp_host_name",
			Type:        proto.ColumnType_STRING,
			Description: "RDP host address",
		},
		{
			Name:        "rdp_host_port",
			Type:        proto.ColumnType_STRING,
			Description: "RDP port (default: 3389)",
		},
		{
			Name:        "rdp_user_groups",
			Type:        proto.ColumnType_STRING,
			Description: "RDP user groups for dynamic users",
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
			Name:        "warn_user_before_expiration",
			Type:        proto.ColumnType_INT,
			Description: "WarnBeforeUserExpiration",
		},
	}
}

func listAkeylessAkeylessDynamicSecretRdp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
