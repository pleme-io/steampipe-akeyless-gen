package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretMongodb() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_mongodb",
		Description: "Manages a MongoDB dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretMongodb,
		},
		Columns: akeylessAkeylessDynamicSecretMongodbColumns(),
	}
}

func akeylessAkeylessDynamicSecretMongodbColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "item_custom_fields",
			Type:        proto.ColumnType_JSON,
			Description: "Additional custom fields to associate with the item",
		},
		{
			Name:        "mongodb_atlas_api_private_key",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB Atlas API private key",
		},
		{
			Name:        "mongodb_atlas_api_public_key",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB Atlas API public key",
		},
		{
			Name:        "mongodb_atlas_project_id",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB Atlas project ID",
		},
		{
			Name:        "mongodb_custom_data",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB custom data",
		},
		{
			Name:        "mongodb_default_auth_db",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB default authentication database",
		},
		{
			Name:        "mongodb_host_port",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB host:port",
		},
		{
			Name:        "mongodb_name",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB database name",
		},
		{
			Name:        "mongodb_password",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB admin password",
		},
		{
			Name:        "mongodb_roles",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB roles (e.g., readWrite@db)",
		},
		{
			Name:        "mongodb_scopes",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB Scopes (Atlas only)",
		},
		{
			Name:        "mongodb_server_uri",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB server URI",
		},
		{
			Name:        "mongodb_uri_options",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB URI options",
		},
		{
			Name:        "mongodb_username",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB admin username",
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

func listAkeylessAkeylessDynamicSecretMongodb(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
