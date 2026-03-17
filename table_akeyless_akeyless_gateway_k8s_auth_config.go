package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessGatewayK8sAuthConfig() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_gateway_k8s_auth_config",
		Description: "Manages a gateway Kubernetes auth config in Akeyless",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessGatewayK8sAuthConfig,
		},
		Columns: akeylessAkeylessGatewayK8sAuthConfigColumns(),
	}
}

func akeylessAkeylessGatewayK8sAuthConfigColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_id",
			Type:        proto.ColumnType_STRING,
			Description: "The access ID of the Kubernetes auth method",
		},
		{
			Name:        "cluster_api_type",
			Type:        proto.ColumnType_STRING,
			Description: "Cluster access type. options: [native_k8s, rancher]",
		},
		{
			Name:        "disable_issuer_validation",
			Type:        proto.ColumnType_STRING,
			Description: "Disable issuer validation [true/false]",
		},
		{
			Name:        "k8s_auth_type",
			Type:        proto.ColumnType_STRING,
			Description: "K8S auth type [token/certificate]. (relevant for \"native_k8s\" only)",
		},
		{
			Name:        "k8s_ca_cert",
			Type:        proto.ColumnType_STRING,
			Description: "The CA Certificate (base64 encoded) to use to call into the kubernetes API server",
		},
		{
			Name:        "k8s_client_certificate",
			Type:        proto.ColumnType_STRING,
			Description: "Content of the k8 client certificate (PEM format) in a Base64 format (relevant for \"native_k8s\" only)",
		},
		{
			Name:        "k8s_client_key",
			Type:        proto.ColumnType_STRING,
			Description: "Content of the k8 client private key (PEM format) in a Base64 format (relevant for \"native_k8s\" only)",
		},
		{
			Name:        "k8s_host",
			Type:        proto.ColumnType_STRING,
			Description: "The URL of the kubernetes API server",
		},
		{
			Name:        "k8s_issuer",
			Type:        proto.ColumnType_STRING,
			Description: "The Kubernetes JWT issuer name. K8SIssuer is the claim that specifies who issued the Kubernetes token",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "K8S Auth config name",
		},
		{
			Name:        "rancher_api_key",
			Type:        proto.ColumnType_STRING,
			Description: "The api key used to access the TokenReview API to validate other JWTs (relevant for \"rancher\" only)",
		},
		{
			Name:        "rancher_cluster_id",
			Type:        proto.ColumnType_STRING,
			Description: "The cluster id as define in rancher (relevant for \"rancher\" only)",
		},
		{
			Name:        "signing_key",
			Type:        proto.ColumnType_STRING,
			Description: "The private key (base64 encoded) associated with the public key defined in the Kubernetes auth",
		},
		{
			Name:        "token_exp",
			Type:        proto.ColumnType_INT,
			Description: "Time in seconds of expiration of the Akeyless Kube Auth Method token",
		},
		{
			Name:        "token_reviewer_jwt",
			Type:        proto.ColumnType_STRING,
			Description: "A Kubernetes service account JWT used to access the TokenReview API to validate other JWTs (relevant for \"native_k8s\" only).\nIf not set, the JWT submitted in the authentication process will be used to access the Kubernetes TokenReview API.",
		},
		{
			Name:        "use_gw_service_account",
			Type:        proto.ColumnType_BOOL,
			Description: "Use the GW's service account",
		},
	}
}

func listAkeylessAkeylessGatewayK8sAuthConfig(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
