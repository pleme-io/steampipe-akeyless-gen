package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessEventForwarderWebhook() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_event_forwarder_webhook",
		Description: "Manages a Webhook event forwarder in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessEventForwarderWebhook,
		},
		Columns: akeylessAkeylessEventForwarderWebhookColumns(),
	}
}

func akeylessAkeylessEventForwarderWebhookColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "auth_methods_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Auth Method Event sources",
		},
		{
			Name:        "auth_token",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded Token string for authentication type Token",
		},
		{
			Name:        "auth_type",
			Type:        proto.ColumnType_STRING,
			Description: "The Webhook authentication type [user-pass, bearer-token, certificate]",
		},
		{
			Name:        "client_cert_data",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded PEM certificate, relevant for certificate auth-type",
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
			Name:        "private_key_data",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded PEM RSA Private Key, relevant for certificate auth-type",
		},
		{
			Name:        "runner_type",
			Type:        proto.ColumnType_STRING,
			Description: "The runner_type field.",
		},
		{
			Name:        "server_certificates",
			Type:        proto.ColumnType_STRING,
			Description: "Base64 encoded PEM certificate of the Webhook",
		},
		{
			Name:        "targets_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Targets Event sources",
		},
		{
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Description: "Webhook URL",
		},
	}
}

func listAkeylessAkeylessEventForwarderWebhook(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
