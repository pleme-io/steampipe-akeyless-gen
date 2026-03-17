package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessDynamicSecretGke() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_dynamic_secret_gke",
		Description: "Manages a Google GKE dynamic secret producer",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessDynamicSecretGke,
		},
		Columns: akeylessAkeylessDynamicSecretGkeColumns(),
	}
}

func akeylessAkeylessDynamicSecretGkeColumns() []*plugin.Column {
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

func listAkeylessAkeylessDynamicSecretGke(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
