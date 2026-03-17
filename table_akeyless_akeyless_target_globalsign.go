package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGlobalsign() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_globalsign",
		Description: "Manages a GlobalSign target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGlobalsign,
		},
		Columns: akeylessAkeylessTargetGlobalsignColumns(),
	}
}

func akeylessAkeylessTargetGlobalsignColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "contact_email",
			Type:        proto.ColumnType_STRING,
			Description: "Contact email",
		},
		{
			Name:        "contact_first_name",
			Type:        proto.ColumnType_STRING,
			Description: "Contact first name",
		},
		{
			Name:        "contact_last_name",
			Type:        proto.ColumnType_STRING,
			Description: "Contact last name",
		},
		{
			Name:        "contact_phone",
			Type:        proto.ColumnType_STRING,
			Description: "Contact phone number",
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
			Name:        "profile_id",
			Type:        proto.ColumnType_STRING,
			Description: "GlobalSign profile ID",
		},
		{
			Name:        "timeout",
			Type:        proto.ColumnType_STRING,
			Description: "Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h.",
		},
	}
}

func listAkeylessAkeylessTargetGlobalsign(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
