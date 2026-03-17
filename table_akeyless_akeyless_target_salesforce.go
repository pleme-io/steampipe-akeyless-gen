package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetSalesforce() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_salesforce",
		Description: "Manages a Salesforce target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetSalesforce,
		},
		Columns: akeylessAkeylessTargetSalesforceColumns(),
	}
}

func akeylessAkeylessTargetSalesforceColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "app_private_key_data",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded PEM of the connected app private key (relevant for JWT auth only)",
		},
		{
			Name:        "auth_flow",
			Type:        proto.ColumnType_STRING,
			Description: "Salesforce auth flow: user-password or jwt-bearer",
		},
		{
			Name:        "ca_cert_data",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded PEM cert to use when uploading a new key to Salesforce",
		},
		{
			Name:        "ca_cert_name",
			Type:        proto.ColumnType_STRING,
			Description: "name of the certificate in Salesforce tenant to use when uploading new key",
		},
		{
			Name:        "client_id",
			Type:        proto.ColumnType_STRING,
			Description: "Salesforce connected app client ID",
		},
		{
			Name:        "client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "Salesforce connected app client secret",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "email",
			Type:        proto.ColumnType_STRING,
			Description: "The email of the user attached to the oauth2 app used for connecting to Salesforce",
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
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Target name",
		},
		{
			Name:        "security_token",
			Type:        proto.ColumnType_STRING,
			Description: "Salesforce security token",
		},
		{
			Name:        "tenant_url",
			Type:        proto.ColumnType_STRING,
			Description: "Salesforce tenant URL",
		},
	}
}

func listAkeylessAkeylessTargetSalesforce(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
