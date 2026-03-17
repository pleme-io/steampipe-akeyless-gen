package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetEks() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_eks",
		Description: "Manages an Amazon EKS target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetEks,
		},
		Columns: akeylessAkeylessTargetEksColumns(),
	}
}

func akeylessAkeylessTargetEksColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "eks_access_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "AWS access key ID for EKS",
		},
		{
			Name:        "eks_cluster_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "EKS cluster CA certificate (PEM)",
		},
		{
			Name:        "eks_cluster_endpoint",
			Type:        proto.ColumnType_STRING,
			Description: "EKS cluster API server URL",
		},
		{
			Name:        "eks_cluster_name",
			Type:        proto.ColumnType_STRING,
			Description: "EKS cluster name",
		},
		{
			Name:        "eks_region",
			Type:        proto.ColumnType_STRING,
			Description: "AWS region for EKS cluster",
		},
		{
			Name:        "eks_secret_access_key",
			Type:        proto.ColumnType_STRING,
			Description: "AWS secret access key for EKS",
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

func listAkeylessAkeylessTargetEks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
