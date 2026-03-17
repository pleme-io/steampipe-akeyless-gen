package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretK8s() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_k8s",
		Description: "Manages a Kubernetes dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretK8s,
		},
		Columns: akeylessAkeylessDynamicSecretK8sColumns(),
	}
}

func akeylessAkeylessDynamicSecretK8sColumns() []*plugin.Column {
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
			Name:        "k8s_allowed_namespaces",
			Type:        proto.ColumnType_STRING,
			Description: "Allowed Kubernetes namespaces (comma-separated)",
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
			Name:        "k8s_namespace",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes namespace",
		},
		{
			Name:        "k8s_predefined_role_name",
			Type:        proto.ColumnType_STRING,
			Description: "The pre-existing Role or ClusterRole name to bind the generated ServiceAccount to (relevant only for k8s-service-account-type=dynamic)",
		},
		{
			Name:        "k8s_predefined_role_type",
			Type:        proto.ColumnType_STRING,
			Description: "Specifies the type of the pre-existing K8S role [Role, ClusterRole] (relevant only for k8s-service-account-type=dynamic)",
		},
		{
			Name:        "k8s_rolebinding_yaml_data",
			Type:        proto.ColumnType_STRING,
			Description: "Content of the yaml in a Base64 format.",
		},
		{
			Name:        "k8s_rolebinding_yaml_def",
			Type:        proto.ColumnType_STRING,
			Description: "Path to yaml file that contains definitions of K8S role and role binding (relevant only for k8s-service-account-type=dynamic)",
		},
		{
			Name:        "k8s_service_account",
			Type:        proto.ColumnType_STRING,
			Description: "Kubernetes service account name",
		},
		{
			Name:        "k8s_service_account_type",
			Type:        proto.ColumnType_STRING,
			Description: "K8S ServiceAccount type [fixed, dynamic].",
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
			Name:        "secure_access_dashboard_url",
			Type:        proto.ColumnType_STRING,
			Description: "The K8s dashboard url",
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
			Name:        "use_gw_service_account",
			Type:        proto.ColumnType_BOOL,
			Description: "Use the GW's service account",
		},
		{
			Name:        "user_ttl",
			Type:        proto.ColumnType_STRING,
			Description: "User TTL (e.g., 60m, 12h)",
		},
	}
}

func listAkeylessAkeylessDynamicSecretK8s(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
