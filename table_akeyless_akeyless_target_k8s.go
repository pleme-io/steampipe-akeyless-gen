package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessTargetK8s() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_target_k8s",
		Description: "Manages a Kubernetes target in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessTargetK8s,
		},
		Columns: akeylessAkeylessTargetK8sColumns(),
	}
}

func akeylessAkeylessTargetK8sColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Target description",
		},
		{
			Name:        "k8s_auth_type",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes auth type: token, certificate, or password",
		},
		{
			Name:        "k8s_client_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes client certificate (PEM)",
		},
		{
			Name:        "k8s_client_key",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes client key (PEM)",
		},
		{
			Name:        "k8s_cluster_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes cluster CA certificate (PEM)",
		},
		{
			Name:        "k8s_cluster_endpoint",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes API server URL",
		},
		{
			Name:        "k8s_cluster_name",
			Type:        proto.ColumnType_STRING,
			Description: "K8S cluster name",
		},
		{
			Name:        "k8s_cluster_token",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes bearer token",
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
			Name:        "use_gw_service_account",
			Type:        proto.ColumnType_BOOL,
			Description: "Use gateway service account for authentication",
		},
	}
}

func listAkeylessAkeylessTargetK8s(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
