package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetGke() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_gke",
		Description: "Manages a Google GKE target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetGke,
		},
		Columns: akeylessAkeylessTargetGkeColumns(),
	}
}

func akeylessAkeylessTargetGkeColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "gke_account_key",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account key JSON (base64)",
		},
		{
			Name:        "gke_cluster_cert",
			Type:        proto.ColumnType_STRING,
			Description: "GKE cluster CA certificate (PEM)",
		},
		{
			Name:        "gke_cluster_endpoint",
			Type:        proto.ColumnType_STRING,
			Description: "GKE cluster API server URL",
		},
		{
			Name:        "gke_cluster_name",
			Type:        proto.ColumnType_STRING,
			Description: "GKE cluster name",
		},
		{
			Name:        "gke_service_account_email",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account email",
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
			Name:        "use_gw_cloud_identity",
			Type:        proto.ColumnType_BOOL,
			Description: "Use gateway cloud identity for authentication",
		},
	}
}

func listAkeylessAkeylessTargetGke(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
