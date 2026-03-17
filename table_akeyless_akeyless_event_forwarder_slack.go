package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessEventForwarderSlack() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_event_forwarder_slack",
		Description: "Manages a Slack event forwarder in Akeyless Vault",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessEventForwarderSlack,
		},
		Columns: akeylessAkeylessEventForwarderSlackColumns(),
	}
}

func akeylessAkeylessEventForwarderSlackColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "auth_methods_event_source_locations",
			Type:        proto.ColumnType_JSON,
			Description: "Auth Method Event sources",
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
			Name:        "url",
			Type:        proto.ColumnType_STRING,
			Description: "Slack Webhook URL",
		},
	}
}

func listAkeylessAkeylessEventForwarderSlack(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
