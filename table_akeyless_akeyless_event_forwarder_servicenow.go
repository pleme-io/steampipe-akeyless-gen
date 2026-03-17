package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessEventForwarderServicenow() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_event_forwarder_servicenow",
		Description: "Manages a ServiceNow event forwarder in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessEventForwarderServicenow,
		},
		Columns: akeylessAkeylessEventForwarderServicenowColumns(),
	}
}

func akeylessAkeylessEventForwarderServicenowColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "admin_name",
			Type:        proto.ColumnType_STRING,
			Description: "Workstation Admin Name",
		},
		{
			Name:        "admin_pwd",
			Type:        proto.ColumnType_STRING,
			Description: "Workstation Admin Password",
		},
		{
			Name:        "app_private_key_base64",
			Type:        proto.ColumnType_STRING,
			Description: "The RSA Private Key to use when connecting with jwt authentication",
		},
		{
			Name:        "auth_methods_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Auth Method Event sources",
		},
		{
			Name:        "auth_type",
			Type:        proto.ColumnType_STRING,
			Description: "The authentication type to use [user-pass/jwt]",
		},
		{
			Name:        "client_id",
			Type:        proto.ColumnType_STRING,
			Description: "The client ID to use when connecting with jwt authentication",
		},
		{
			Name:        "client_secret",
			Type:        proto.ColumnType_STRING,
			Description: "The client secret to use when connecting with jwt authentication",
		},
		{
			Name:        "description",
			Type:        proto.ColumnType_STRING,
			Description: "Description of the object",
		},
		{
			Name:        "event_types",
			Type:        proto.ColumnType_JSON,
			Description: "List of event types to notify about [request-access, certificate-pending-expiration, certificate-expired, certificate-provisioning-success, certificate-provisioning-failure, auth-method-pending-expiration, auth-method-expired, next-automatic-rotation, rotated-secret-success, rotated-secret-failure, dynamic-secret-failure, multi-auth-failure, uid-rotation-failure, apply-justification, email-auth-method-approved, usage, rotation-usage, gateway-inactive, static-secret-updated, rate-limiting, usage-report, secret-sync]",
		},
		{
			Name:        "every",
			Type:        proto.ColumnType_STRING,
			Description: "Rate of periodic runner repetition in hours",
		},
		{
			Name:        "gateways_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Event sources",
		},
		{
			Name:        "host",
			Type:        proto.ColumnType_STRING,
			Description: "Workstation Host",
		},
		{
			Name:        "items_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Items Event sources",
		},
		{
			Name:        "key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of a key that used to encrypt the EventForwarder secret value (if empty, the account default protectionKey key will be used)",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "EventForwarder name",
		},
		{
			Name:        "runner_type",
			Type:        proto.ColumnType_STRING,
			Description: "The runner_type field.",
		},
		{
			Name:        "targets_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Targets Event sources",
		},
		{
			Name:        "user_email",
			Type:        proto.ColumnType_STRING,
			Description: "The user email to identify with when connecting with jwt authentication",
		},
	}
}

func listAkeylessAkeylessEventForwarderServicenow(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
