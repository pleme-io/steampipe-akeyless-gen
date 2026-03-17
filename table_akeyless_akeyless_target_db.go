package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetDb() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_db",
		Description: "Manages a database target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetDb,
		},
		Columns: akeylessAkeylessTargetDbColumns(),
	}
}

func akeylessAkeylessTargetDbColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "azure_client_id",
			Type:        proto.ColumnType_STRING,
			Description: "(Optional) Client id (relevant for \"cloud-service-provider\" only)",
		},
		{
			Name:        "azure_client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "(Optional) Client secret (relevant for \"cloud-service-provider\" only)",
		},
		{
			Name:        "azure_tenant_id",
			Type:        proto.ColumnType_STRING,
			Description: "(Optional) Tenant id (relevant for \"cloud-service-provider\" only)",
		},
		{
			Name:        "cloud_service_provider",
			Type:        proto.ColumnType_STRING,
			Description: "(Optional) Cloud service provider (currently only supports Azure)",
		},
		{
			Name:        "cluster_mode",
			Type:        proto.ColumnType_BOOL,
			Description: "Cluster Mode",
		},
		{
			Name:        "comment",
			Type:        proto.ColumnType_STRING,
			Description: "Deprecated - use description",
		},
		{
			Name:        "connection_type",
			Type:        proto.ColumnType_STRING,
			Description: "Type of connection to mssql database [credentials/cloud-identity/wallet/parent-target]",
		},
		{
			Name:        "db_name",
			Type:        proto.ColumnType_STRING,
			Description: "Database name",
		},
		{
			Name:        "db_server_certificates",
			Type:        proto.ColumnType_STRING,
			Description: "Database server TLS certificates (PEM)",
		},
		{
			Name:        "db_server_name",
			Type:        proto.ColumnType_STRING,
			Description: "Database server name for TLS verification",
		},
		{
			Name:        "db_type",
			Type:        proto.ColumnType_STRING,
			Description: "Database type: mysql, mssql, postgresql, mongodb, snowflake, oracle, cassandra, redshift, hanadb",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "host",
			Type:        proto.ColumnType_STRING,
			Description: "Database hostname",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "max_versions",
			Type:        proto.ColumnType_STRING,
			Description: "Set the maximum number of versions, limited by the account settings defaults.",
		},
		{
			Name:        "mongodb_atlas",
			Type:        proto.ColumnType_BOOL,
			Description: "Use MongoDB Atlas connection",
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
			Name:        "mongodb_default_auth_db",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB default authentication database",
		},
		{
			Name:        "mongodb_uri_options",
			Type:        proto.ColumnType_STRING,
			Description: "MongoDB additional URI options",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "oracle_service_name",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle service name",
		},
		{
			Name:        "oracle_wallet_login_type",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle Wallet login type (password/mtls)",
		},
		{
			Name:        "oracle_wallet_p12_file_data",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle wallet p12 file data in base64",
		},
		{
			Name:        "oracle_wallet_sso_file_data",
			Type:        proto.ColumnType_STRING,
			Description: "Oracle wallet sso file data in base64",
		},
		{
			Name:        "parent_target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Name of the parent target, relevant only when connection-type is parent-target",
		},
		{
			Name:        "port",
			Type:        proto.ColumnType_STRING,
			Description: "Database port",
		},
		{
			Name:        "pwd",
			Type:        proto.ColumnType_STRING,
			Description: "Database password",
		},
		{
			Name:        "snowflake_account",
			Type:        proto.ColumnType_STRING,
			Description: "Snowflake account identifier",
		},
		{
			Name:        "snowflake_api_private_key",
			Type:        proto.ColumnType_STRING,
			Description: "RSA Private key (base64 encoded)",
		},
		{
			Name:        "snowflake_api_private_key_password",
			Type:        proto.ColumnType_STRING,
			Description: "The Private key passphrase",
		},
		{
			Name:        "ssl",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable SSL connection",
		},
		{
			Name:        "ssl_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "Client SSL certificate (PEM)",
		},
		{
			Name:        "user_name",
			Type:        proto.ColumnType_STRING,
			Description: "Database username",
		},
	}
}

func listAkeylessAkeylessTargetDb(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
