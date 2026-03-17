package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretSnowflake() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_snowflake",
		Description: "Manages a Snowflake dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretSnowflake,
		},
		Columns: akeylessAkeylessDynamicSecretSnowflakeColumns(),
	}
}

func akeylessAkeylessDynamicSecretSnowflakeColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "account",
			Type:        proto.ColumnType_STRING,
			Description: "Account name",
		},
		{
			Name:        "account_password",
			Type:        proto.ColumnType_STRING,
			Description: "Database Password",
		},
		{
			Name:        "account_username",
			Type:        proto.ColumnType_STRING,
			Description: "Database Username",
		},
		{
			Name:        "auth_mode",
			Type:        proto.ColumnType_STRING,
			Description: "The authentication mode for the temporary user [password/key]",
		},
		{
			Name:        "custom_username_template",
			Type:        proto.ColumnType_STRING,
			Description: "Customize how temporary usernames are generated using go template",
		},
		{
			Name:        "db_name",
			Type:        proto.ColumnType_STRING,
			Description: "Database name",
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
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "key_algo",
			Type:        proto.ColumnType_STRING,
			Description: "The key_algo field.",
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
			Name:        "private_key",
			Type:        proto.ColumnType_STRING,
			Description: "RSA Private key (base64 encoded)",
		},
		{
			Name:        "private_key_passphrase",
			Type:        proto.ColumnType_STRING,
			Description: "The Private key passphrase",
		},
		{
			Name:        "role",
			Type:        proto.ColumnType_STRING,
			Description: "User role",
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
			Name:        "warehouse",
			Type:        proto.ColumnType_STRING,
			Description: "Warehouse name",
		},
	}
}

func listAkeylessAkeylessDynamicSecretSnowflake(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
