package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretAzure() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_azure",
		Description: "Manages an Azure dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretAzure,
		},
		Columns: akeylessAkeylessDynamicSecretAzureColumns(),
	}
}

func akeylessAkeylessDynamicSecretAzureColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "app_obj_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure application object ID",
		},
		{
			Name:        "azure_administrative_unit",
			Type:        proto.ColumnType_STRING,
			Description: "Azure AD administrative unit (relevant only when azure-user-portal-access=true)",
		},
		{
			Name:        "azure_client_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure client/application ID",
		},
		{
			Name:        "azure_client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "Azure client secret",
		},
		{
			Name:        "azure_tenant_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure tenant ID",
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
			Description: "FixedUserClaimKeyname",
		},
		{
			Name:        "fixed_user_only",
			Type:        proto.ColumnType_BOOL,
			Description: "Fixed user",
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
			Name:        "user_group_obj_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure user group object ID",
		},
		{
			Name:        "user_portal_access",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable Azure portal access",
		},
		{
			Name:        "user_principal_name",
			Type:        proto.ColumnType_STRING,
			Description: "User Principal Name",
		},
		{
			Name:        "user_programmatic_access",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable Azure programmatic access",
		},
		{
			Name:        "user_role_template_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure role template ID",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretAzure(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
