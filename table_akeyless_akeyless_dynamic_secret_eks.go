package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretEks() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_eks",
		Description: "Manages an Amazon EKS dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretEks,
		},
		Columns: akeylessAkeylessDynamicSecretEksColumns(),
	}
}

func akeylessAkeylessDynamicSecretEksColumns() []*plugin.Column {
	return []*plugin.Column{
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
			Name:        "eks_access_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "AWS access key ID for EKS",
		},
		{
			Name:        "eks_assume_role",
			Type:        proto.ColumnType_STRING,
			Description: "AWS role ARN to assume for EKS",
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
			Name:        "producer_encryption_key_name",
			Type:        proto.ColumnType_STRING,
			Description: "Dynamic producer encryption key",
		},
		{
			Name:        "secure_access_allow_port_forwading",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable Port forwarding while using CLI access",
		},
		{
			Name:        "secure_access_cluster_endpoint",
			Type:        proto.ColumnType_STRING,
			Description: "The K8s cluster endpoint URL",
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

func listAkeylessAkeylessDynamicSecretEks(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
