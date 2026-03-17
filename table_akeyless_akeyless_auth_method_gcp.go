package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessAuthMethodGcp() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_auth_method_gcp",
		Description: "Manages a GCP authentication method",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessAuthMethodGcp,
		},
		Columns: akeylessAkeylessAuthMethodGcpColumns(),
	}
}

func akeylessAkeylessAuthMethodGcpColumns() []*plugin.Column {
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
			Description: "GCP audience for IAM auth",
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
			Name:        "bound_labels",
			Type:        proto.ColumnType_JSON,
			Description: "GCP labels to restrict access (GCE type)",
		},
		{
			Name:        "bound_projects",
			Type:        proto.ColumnType_JSON,
			Description: "GCP project IDs to restrict access",
		},
		{
			Name:        "bound_regions",
			Type:        proto.ColumnType_JSON,
			Description: "GCP regions to restrict access (GCE type)",
		},
		{
			Name:        "bound_service_accounts",
			Type:        proto.ColumnType_JSON,
			Description: "GCP service account emails to restrict access",
		},
		{
			Name:        "bound_zones",
			Type:        proto.ColumnType_JSON,
			Description: "GCP zones to restrict access (GCE type)",
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
			Name:        "service_account_creds_data",
			Type:        proto.ColumnType_STRING,
			Description: "GCP service account credentials JSON",
		},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "GCP auth type: gce or iam",
		},
		{
			Name:        "unique_identifier",
			Type:        proto.ColumnType_STRING,
			Description: "A unique identifier (ID) value which is a \"sub claim\" name that contains details uniquely identifying that resource. This \"sub claim\" is used to distinguish between different identities.",
		},
	}
}

func listAkeylessAkeylessAuthMethodGcp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
