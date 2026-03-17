package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodK8s() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_k8s",
		Description: "Manages a Kubernetes authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodK8s,
		},
		Columns: akeylessAkeylessAuthMethodK8sColumns(),
	}
}

func akeylessAkeylessAuthMethodK8sColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "access_expires",
			Type:        proto.ColumnType_INT,
			Description: "Access expiration in Unix time (0 = no expiry)",
		},
		{
			Name:        "allowed_client_type",
			Type:        proto.ColumnType_JSON,
			Description: "limit the auth method usage for specific client types [cli,ui,gateway-admin,sdk,mobile,extension]",
		},
		{
			Name:        "audience",
			Type:        proto.ColumnType_STRING,
			Description: "The audience in the Kubernetes JWT that the access is restricted to",
		},
		{
			Name:        "audit_logs_claims",
			Type:        proto.ColumnType_JSON,
			Description: "Subclaims to include in audit logs, e.g \"--audit-logs-claims email --audit-logs-claims username\"",
		},
		{
			Name:        "bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "CIDR whitelist for access",
		},
		{
			Name:        "bound_namespaces",
			Type:        proto.ColumnType_JSON,
			Description: "Kubernetes namespaces to restrict access",
		},
		{
			Name:        "bound_pod_names",
			Type:        proto.ColumnType_JSON,
			Description: "Kubernetes pod names to restrict access",
		},
		{
			Name:        "bound_sa_names",
			Type:        proto.ColumnType_JSON,
			Description: "A list of service account names that the access is restricted to",
		},
		{
			Name:        "delete_protection",
			Type:        proto.ColumnType_BOOL,
			Description: "Enable delete protection",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Auth Method description",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the auth method would you like to be notified.",
		},
		{
			Name:        "force_sub_claims",
			Type:        proto.ColumnType_BOOL,
			Description: "Force sub-claims enforcement",
		},
		{
			Name:        "gen_key",
			Type:        proto.ColumnType_STRING,
			Description: "Automatically generate a key pair (true/false)",
		},
		{
			Name:        "gw_bound_ips",
			Type:        proto.ColumnType_JSON,
			Description: "Gateway CIDR whitelist",
		},
		{
			Name:        "jwt_ttl",
			Type:        proto.ColumnType_INT,
			Description: "JWT TTL in seconds",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Auth Method name",
		},
		{
			Name:        "product_type",
			Type:        proto.ColumnType_JSON,
			Description: "Choose the relevant product type for the auth method [sm, sra, pm, dp, ca]",
		},
		{
			Name:        "public_key",
			Type:        proto.ColumnType_STRING,
			Description: "Public key PEM for token verification",
		},
	}
}

func listAkeylessAkeylessAuthMethodK8s(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
