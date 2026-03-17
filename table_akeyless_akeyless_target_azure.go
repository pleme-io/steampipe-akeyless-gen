package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetAzure() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_azure",
		Description: "Manages an Azure target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetAzure,
		},
		Columns: akeylessAkeylessTargetAzureColumns(),
	}
}

func akeylessAkeylessTargetAzureColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "azure_cloud",
			Type:        proto.ColumnType_STRING,
			Description: "Azure cloud environment to use. Values: AzureCloud (default), AzureUSGovernment, AzureChinaCloud.",
		},
		{
			Name:        "client_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure client/application ID",
		},
		{
			Name:        "client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "Azure client secret",
		},
		{
			Name:        "connection_type",
			Type:        proto.ColumnType_STRING,
			Description: "Type of connection [credentials/cloud-identity]",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
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
			Name:        "resource_group_name",
			Type:        proto.ColumnType_STRING,
			Description: "Azure resource group name",
		},
		{
			Name:        "resource_name",
			Type:        proto.ColumnType_STRING,
			Description: "Azure resource name",
		},
		{
			Name:        "subscription_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure subscription ID",
		},
		{
			Name:        "tenant_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure tenant ID",
		},
		{
			Name:        "use_gw_cloud_identity",
			Type:        proto.ColumnType_BOOL,
			Description: "Use gateway cloud identity for authentication",
		},
	}
}

func listAkeylessAkeylessTargetAzure(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
