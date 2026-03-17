package akeyless

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableAkeylessAkeylessGatewayMigration() *plugin.Table {
	return &plugin.Table{
		Name:        "akeyless_akeyless_gateway_migration",
		Description: "Manages a gateway migration in Akeyless",
		List: &plugin.ListConfig{
			Hydrate: listAkeylessAkeylessGatewayMigration,
		},
		Columns: akeylessAkeylessGatewayMigrationColumns(),
	}
}

func akeylessAkeylessGatewayMigrationColumns() []*plugin.Column {
	return []*plugin.Column{
		{
			Name:        "ServiceAccountKeyDecoded",
			Type:        proto.ColumnType_STRING,
			Description: "The ServiceAccountKeyDecoded field.",
		},
		{
			Name:        "ad_auto_rotate",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable automatic/recurrent rotation for migrated secrets. Default is false: only manual rotation is allowed for migrated secrets. If set to true, this command should be combined with --ad-rotation-interval and --ad-rotation-hour parameters (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_cert_expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of discovered certificates would you like to be notified (Relevant only for Active Directory migration with certificate discovery enabled)",
		},
		{
			Name:        "ad_certificates_path_template",
			Type:        proto.ColumnType_STRING,
			Description: "Path location template for migrating certificates e.g.: /Certificates/{{COMMON_NAME}} (Relevant only for Active Directory migration with certificate discovery enabled)",
		},
		{
			Name:        "ad_computer_base_dn",
			Type:        proto.ColumnType_STRING,
			Description: "Distinguished Name of Computer objects (servers) to search in Active Directory e.g.: CN=Computers,DC=example,DC=com (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_discover_iis_app",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable discovery of IIS application from each domain server as part of the SSH/Windows Rotated Secrets. Default is false. (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_discover_services",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable discovery of Windows services from each domain server as part of the SSH/Windows Rotated Secrets. Default is false. (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_discovery_types",
			Type:        proto.ColumnType_JSON,
			Description: "Set migration discovery types (domain-users, computers, local-users). (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_domain_name",
			Type:        proto.ColumnType_STRING,
			Description: "Active Directory Domain Name (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_domain_users_path_template",
			Type:        proto.ColumnType_STRING,
			Description: "Path location template for migrating domain users as Rotated Secrets e.g.: .../DomainUsers/{{USERNAME}} (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_local_users_ignore",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of Local Users which should not be migrated (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_local_users_path_template",
			Type:        proto.ColumnType_STRING,
			Description: "Path location template for migrating domain users as Rotated Secrets e.g.: .../LocalUsers/{{COMPUTER_NAME}}/{{USERNAME}} (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_os_filter",
			Type:        proto.ColumnType_STRING,
			Description: "Filter by Operating System to run the migration, can be used with wildcards, e.g. SRV20* (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_rotation_hour",
			Type:        proto.ColumnType_INT,
			Description: "The hour of the scheduled rotation in UTC (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_rotation_interval",
			Type:        proto.ColumnType_INT,
			Description: "The number of days to wait between every automatic rotation [1-365] (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_sra_enable_rdp",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable RDP Secure Remote Access for the migrated local users rotated secrets. Default is false: rotated secrets will not be created with SRA (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_ssh_port",
			Type:        proto.ColumnType_STRING,
			Description: "Set the SSH Port for further connection to the domain servers. Default is port 22 (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_target_format",
			Type:        proto.ColumnType_STRING,
			Description: "Relevant only for ad-discovery-types=computers. For linked, all computers will be migrated into a linked target(s). if set with regular, the migration will create a target for each computer.",
		},
		{
			Name:        "ad_target_name",
			Type:        proto.ColumnType_STRING,
			Description: "Active Directory LDAP Target Name. Server type should be Active Directory (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_targets_path_template",
			Type:        proto.ColumnType_STRING,
			Description: "Path location template for migrating domain servers as SSH/Windows Targets e.g.: .../Servers/{{COMPUTER_NAME}} (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_targets_type",
			Type:        proto.ColumnType_STRING,
			Description: "Set the target type of the domain servers [ssh/windows](Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_user_base_dn",
			Type:        proto.ColumnType_STRING,
			Description: "Distinguished Name of User objects to search in Active Directory, e.g.: CN=Users,DC=example,DC=com (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_user_groups",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of domain groups from which privileged domain users will be migrated. If empty, migrate all users based on the --ad-user-base-dn (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_winrm_over_http",
			Type:        proto.ColumnType_STRING,
			Description: "Use WinRM over HTTP, by default runs over HTTPS",
		},
		{
			Name:        "ad_winrm_port",
			Type:        proto.ColumnType_STRING,
			Description: "Set the WinRM Port for further connection to the domain servers. Default is 5986 (Relevant only for Active Directory migration)",
		},
		{
			Name:        "ad_discover_local_users",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable discovery of local users from each domain server and migrate them as SSH/Windows Rotated Secrets. Default is false: only domain users will be migrated. Discovery of local users might require further installation of SSH on the servers, based on the supplied computer base DN. This will be implemented automatically as part of the migration process (Relevant only for Active Directory migration)\nDeprecated: use AdDiscoverTypes",
		},
		{
			Name:        "ai_certificate_discovery",
			Type:        proto.ColumnType_STRING,
			Description: "Enable AI-assisted certificate discovery (only when AI Insight is enabled on the Gateway)",
		},
		{
			Name:        "aws_key",
			Type:        proto.ColumnType_STRING,
			Description: "AWS Secret Access Key (relevant only for AWS migration)",
		},
		{
			Name:        "aws_key_id",
			Type:        proto.ColumnType_STRING,
			Description: "AWS Access Key ID with sufficient permissions to get all secrets, e.g. 'arn:aws:secretsmanager:[Region]:[AccountId]:secret:[/path/to/secrets/*]' (relevant only for AWS migration)",
		},
		{
			Name:        "aws_region",
			Type:        proto.ColumnType_STRING,
			Description: "AWS region of the required Secrets Manager (relevant only for AWS migration)",
		},
		{
			Name:        "azure_client_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure Key Vault Access client ID, should be Azure AD App with a service principal (relevant only for Azure Key Vault migration)",
		},
		{
			Name:        "azure_kv_name",
			Type:        proto.ColumnType_STRING,
			Description: "Azure Key Vault Name (relevant only for Azure Key Vault migration)",
		},
		{
			Name:        "azure_secret",
			Type:        proto.ColumnType_STRING,
			Description: "Azure Key Vault secret (relevant only for Azure Key Vault migration)",
		},
		{
			Name:        "azure_tenant_id",
			Type:        proto.ColumnType_STRING,
			Description: "Azure Key Vault Access tenant ID (relevant only for Azure Key Vault migration)",
		},
		{
			Name:        "conjur_account",
			Type:        proto.ColumnType_STRING,
			Description: "Conjur account name set on your Conjur server (relevant only for Conjur migration).",
		},
		{
			Name:        "conjur_api_key",
			Type:        proto.ColumnType_STRING,
			Description: "Conjur API Key for the specified user (relevant only for Conjur migration).",
		},
		{
			Name:        "conjur_url",
			Type:        proto.ColumnType_STRING,
			Description: "Conjur server base URL (relevant only for Conjur migration).\nIf conjur-url is HTTPS and Conjur uses a private CA/self-signed certificate,\nmake the CA bundle available on the Gateway and set CONJUR_SSL_CERT_PATH to its path.",
		},
		{
			Name:        "conjur_username",
			Type:        proto.ColumnType_STRING,
			Description: "Conjur username used to authenticate (relevant only for Conjur migration).",
		},
		{
			Name:        "expiration_event_in",
			Type:        proto.ColumnType_JSON,
			Description: "How many days before the expiration of the certificate would you like to be notified.",
		},
		{
			Name:        "gcp_key",
			Type:        proto.ColumnType_STRING,
			Description: "Base64-encoded GCP Service Account private key text with sufficient permissions to Secrets Manager, Minimum required permission is Secret Manager Secret Accessor, e.g. 'roles/secretmanager.secretAccessor' (relevant only for GCP migration)",
		},
		{
			Name:        "gcp_project_id",
			Type:        proto.ColumnType_STRING,
			Description: "GCP Project ID (cross-project override)",
		},
		{
			Name:        "hashi_json",
			Type:        proto.ColumnType_STRING,
			Description: "Import secret key as json value or independent secrets (relevant only for HasiCorp Vault migration) [true/false]",
		},
		{
			Name:        "hashi_ns",
			Type:        proto.ColumnType_JSON,
			Description: "HashiCorp Vault Namespaces is a comma-separated list of namespaces which need to be imported into Akeyless Vault. For every provided namespace, all its child namespaces are imported as well, e.g. nmsp/subnmsp1/subnmsp2,nmsp/anothernmsp. By default, import all namespaces (relevant only for HasiCorp Vault migration)",
		},
		{
			Name:        "hashi_token",
			Type:        proto.ColumnType_STRING,
			Description: "HashiCorp Vault access token with sufficient permissions to preform list & read operations on secrets objects (relevant only for HasiCorp Vault migration)",
		},
		{
			Name:        "hashi_url",
			Type:        proto.ColumnType_STRING,
			Description: "HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 (relevant only for HasiCorp Vault migration)",
		},
		{
			Name:        "hosts",
			Type:        proto.ColumnType_STRING,
			Description: "A comma separated list of IPs, CIDR ranges, or DNS names to scan",
		},
		{
			Name:        "k8s_ca_certificate",
			Type:        proto.ColumnType_JSON,
			Description: "For Certificate Authentication method\nK8s Cluster CA certificate (relevant only for K8s migration with Certificate Authentication method)",
		},
		{
			Name:        "k8s_client_certificate",
			Type:        proto.ColumnType_JSON,
			Description: "K8s Client certificate with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Certificate Authentication method)",
		},
		{
			Name:        "k8s_client_key",
			Type:        proto.ColumnType_JSON,
			Description: "K8s Client key (relevant only for K8s migration with Certificate Authentication method)",
		},
		{
			Name:        "k8s_namespace",
			Type:        proto.ColumnType_STRING,
			Description: "K8s Namespace, Use this field to import secrets from a particular namespace only. By default, the secrets are imported from all namespaces (relevant only for K8s migration)",
		},
		{
			Name:        "k8s_password",
			Type:        proto.ColumnType_STRING,
			Description: "K8s Client password (relevant only for K8s migration with Password Authentication method)",
		},
		{
			Name:        "k8s_skip_system",
			Type:        proto.ColumnType_BOOL,
			Description: "K8s Skip Control Plane Secrets, This option allows to avoid importing secrets from system namespaces (relevant only for K8s migration)",
		},
		{
			Name:        "k8s_token",
			Type:        proto.ColumnType_STRING,
			Description: "For Token Authentication method\nK8s Bearer Token with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Token Authentication method)",
		},
		{
			Name:        "k8s_url",
			Type:        proto.ColumnType_STRING,
			Description: "K8s API Server URL, e.g. https://k8s-api.mycompany.com:6443 (relevant only for K8s migration)",
		},
		{
			Name:        "k8s_username",
			Type:        proto.ColumnType_STRING,
			Description: "For Password Authentication method\nK8s Client username with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Password Authentication method)",
		},
		{
			Name:        "name",
			Type:        proto.ColumnType_STRING,
			Description: "Migration name",
		},
		{
			Name:        "port_ranges",
			Type:        proto.ColumnType_STRING,
			Description: "A comma separated list of port ranges\nExamples: \"80,443\" or \"80,443,8080-8090\" or \"443\"",
		},
		{
			Name:        "protection_key",
			Type:        proto.ColumnType_STRING,
			Description: "The name of the key that protects the classic key value (if empty, the account default key will be used)",
		},
		{
			Name:        "si_auto_rotate",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable automatic/recurrent rotation for migrated secrets. Default is false: only manual rotation is allowed for migrated secrets. If set to true, this command should be combined with --si-rotation-interval and --si-rotation-hour parameters (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_rotation_hour",
			Type:        proto.ColumnType_INT,
			Description: "The hour of the scheduled rotation in UTC (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_rotation_interval",
			Type:        proto.ColumnType_INT,
			Description: "The number of days to wait between every automatic rotation [1-365] (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_sra_enable_rdp",
			Type:        proto.ColumnType_STRING,
			Description: "Enable/Disable RDP Secure Remote Access for the migrated local users rotated secrets. Default is false: rotated secrets will not be created with SRA (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_target_name",
			Type:        proto.ColumnType_STRING,
			Description: "SSH, Windows or Linked Target Name. (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_user_groups",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of groups to migrate users from. If empty, all users from all groups will be migrated (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_users_ignore",
			Type:        proto.ColumnType_STRING,
			Description: "Comma-separated list of Local Users which should not be migrated (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "si_users_path_template",
			Type:        proto.ColumnType_STRING,
			Description: "Path location template for migrating users as Rotated Secrets e.g.: .../Users/{{COMPUTER_NAME}}/{{USERNAME}} (Relevant only for Server Inventory migration)",
		},
		{
			Name:        "target_location",
			Type:        proto.ColumnType_STRING,
			Description: "Target location in Akeyless for imported secrets",
		},
		{
			Name:        "type",
			Type:        proto.ColumnType_STRING,
			Description: "Migration type (hashi/aws/gcp/k8s/azure_kv/conjur/active_directory/server_inventory/certificate)",
		},
		{
			Name:        "use_gw_cloud_identity",
			Type:        proto.ColumnType_BOOL,
			Description: "Use the GW's Cloud IAM",
		},
	}
}

func listAkeylessAkeylessGatewayMigration(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// TODO: Implement list hydrate function.
	return nil, nil
}
