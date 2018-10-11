package cli

import (
	"fmt"
	"reflect"

	"github.com/hortonworks/cb-cli/cli/utils"
	"github.com/urfave/cli"
)

var REQUIRED = RequiredFlag{true}
var OPTIONAL = RequiredFlag{false}

var (
	FlDebugOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},
	}
	FlWaitOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "wait",
			Usage: "wait for the operation to finish, no argument required",
		},
	}
	FlInputJson = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "cli-input-json",
			Usage: "user provided file with json content",
		},
	}
	FlOutputOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "output",
			Usage:  "supported formats: json, yaml, table (default: \"json\")",
			EnvVar: "CB_OUT_FORMAT",
		},
	}
	FlProfileOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "profile",
			Usage:  "selects a config profile to use",
			EnvVar: "CB_PROFILE",
		},
	}
	FlAuthTypeOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "auth-type",
			Usage:  "authentication method to use, values: [" + OAUTH2 + ", " + BASIC + "]",
			EnvVar: "CB_AUTH_TYPE",
		},
	}
	FlForceOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "force",
			Usage: "force the operation",
		},
	}
	FlServerOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "server",
			Usage:  "server address",
			EnvVar: "CB_SERVER_ADDRESS",
		},
	}
	FlServerRequired = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:   "server",
			Usage:  "server address",
			EnvVar: "CB_SERVER_ADDRESS",
		},
	}
	FlUsername = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "username",
			Usage:  "user name (e-mail address)",
			EnvVar: "CB_USER_NAME",
		},
	}
	FlUsernameRequired = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:   "username",
			Usage:  "user name (e-mail address)",
			EnvVar: "CB_USER_NAME",
		},
	}
	FlPassword = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:   "password",
			Usage:  "password",
			EnvVar: "CB_PASSWORD",
		},
	}
	FlName = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "name",
			Usage: "name of resource",
		},
	}
	FlNameOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "name",
			Usage: "name of resource",
		},
	}
	FlDescriptionOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "description",
			Usage: "description of resource",
		},
	}
	FlPublicOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "public",
			Usage: "public in account",
		},
	}
	FlRoleARN = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "role-arn",
		},
	}
	FlAccessKey = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "access-key",
		},
	}
	FlSecretKey = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "secret-key",
		},
	}
	FlProjectId = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "project-id",
		},
	}
	FlServiceAccountId = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "service-account-id",
		},
	}
	FlServiceAccountPrivateKeyFile = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "service-account-private-key-file",
		},
	}
	FlServiceAccountJsonFile = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "service-account-json-file",
		},
	}
	FlTenantUser = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "tenant-user",
		},
	}
	FlTenantPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "tenant-password",
		},
	}
	FlTenantName = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "tenant-name",
		},
	}
	FlEndpoint = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "endpoint",
		},
	}
	FlKeystoneScopeOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name: "keystone-scope",
		},
	}
	FlUserDomain = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "user-domain",
		},
	}
	FlProjectDomainNameOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name: "project-domain-name",
		},
	}
	FlDomainNameOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name: "domain-name",
		},
	}
	FlProjectNameOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name: "project-name",
		},
	}
	FlFacingOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name: "facing",
		},
	}
	FlSubscriptionId = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "subscription-id",
		},
	}
	FlTenantId = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "tenant-id",
		},
	}
	FlAppId = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "app-id",
		},
	}
	FlAppPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name: "app-password",
		},
	}
	FlFile = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "file",
			Usage: "location of the Ambari blueprint JSON file",
		},
	}
	FlURL = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "url",
			Usage: "URL location of the JSON file",
		},
	}
	FlBlueprintName = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "blueprint-name",
			Usage: "name of the blueprint",
		},
	}
	FlBlueprintNameOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "blueprint-name",
			Usage: "name of the blueprint",
		},
	}
	FlBlueprintFileOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "blueprint-file",
			Usage: "location of the blueprint file",
		},
	}
	FlCloudStorageTypeOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "cloud-storage",
			Usage: "type of the cloud storage [wasb/WASB, adls/ADLS, s3/S3, gcs/GCS]",
		},
	}
	FlExecutionType = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "execution-type",
			Usage: "type of execution [pre-ambari-start, pre-termination, post-ambari-start, post-cluster-install]",
		},
	}
	FlAmbariPasswordOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "input-json-param-password",
			Usage: "password of the cluster and ambari",
		},
	}
	FlGroupName = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "group-name",
			Usage: "name of the group to scale",
		},
	}
	FlDesiredNodeCount = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "desired-node-count",
			Usage: "desired number of nodes",
		},
	}
	FlOldPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "old-password",
			Usage: "old password of ambari",
		},
	}
	FlNewPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "new-password",
			Usage: "new password of ambari",
		},
	}
	FlAmbariUser = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ambari-user",
			Usage: "user of ambari",
		},
	}
	FlLdapServer = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-server",
			Usage: "address of the ldap server (e.g: ldap://10.0.0.1:384)",
		},
	}
	FlLdapSecureOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "ldaps",
			Usage: "set ldaps if the ldap is secured with SSL",
		},
	}
	FlLdapDomain = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-domain",
			Usage: "ldap domain (e.g: ad.cb.com)",
		},
	}
	FlLdapBindDN = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-bind-dn",
			Usage: "ldap bind dn (e.g: CN=Administrator,CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapUserNameAttribute = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-name-attribute",
			Usage: "ldap user name attribute",
		},
	}
	FlLdapUserObjectClass = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-object-class",
			Usage: "ldap user object class",
		},
	}
	FlLdapGroupMemberAttribute = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-member-attribute",
			Usage: "ldap group member attribute",
		},
	}
	FlLdapGroupNameAttribute = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-name-attribute",
			Usage: "ldap group name attribute",
		},
	}
	FlLdapGroupObjectClass = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-object-class",
			Usage: "ldap group object class",
		},
	}
	FlLdapBindPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-bind-password",
			Usage: "ldap bind password",
		},
	}
	FlLdapDirectoryType = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-directory-type",
			Usage: "ldap directory type (LDAP or ACTIVE_DIRECTORY)",
		},
	}
	FlLdapUserSearchBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-search-base",
			Usage: "ldap user search base (e.g: CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapUserDnPattern = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-dn-pattern",
			Usage: "ldap userDnPattern (e.g: CN={0},DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapGroupSearchBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-search-base",
			Usage: "ldap group search base (e.g: OU=scopes,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapAdminGroup = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "ldap-admin-group",
			Usage: "ldap group of administrators",
		},
	}
	FlLdapUserToCreate = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-create",
			Usage: "name of the ldap user (e.g user will create CN=user)",
		},
	}
	FlLdapUserToCreatePassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-create-password",
			Usage: "password of the user",
		},
	}
	FlLdapUserToCreateEmail = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-create-email",
			Usage: "email of the ldap user (it will set the mail attribute)",
		},
	}
	FlLdapUserToCreateBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-create-base",
			Usage: "base DN where the user will be created (e.g: CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapUserToCreateGroups = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-create-groups",
			Usage: "semicolon separated list of group DNs that the user will be added to (e.g: OU=cloudbreak,CN=Users,DC=ad,DC=cb,DC=com;)",
		},
	}
	FlLdapGroupToCreate = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-to-create",
			Usage: "name of the ldap group (e.g cloudbreak will create CN=cloudbreak)",
		},
	}
	FlLdapGroupToCreateBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-to-create-base",
			Usage: "base DN where the group will be created (e.g: CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapGroupToDelete = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-to-delete",
			Usage: "name of the ldap group",
		},
	}
	FlLdapGroupToDeleteBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-group-to-delete-base",
			Usage: "base DN from where the group will be deleted (e.g: CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlLdapUserToDelete = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-delete",
			Usage: "name of the ldap user to delete",
		},
	}
	FlLdapUserToDeleteBase = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "ldap-user-to-delete-base",
			Usage: "base DN from where the user will be deleted (e.g: CN=Users,DC=ad,DC=cb,DC=com)",
		},
	}
	FlKerberosPasswordOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "kerberos-password",
			Usage: "kerberos password",
		},
	}
	FlKerberosPrincipalOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "kerberos-principal",
			Usage: "kerberos principal",
		},
	}
	FlCredential = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "credential",
			Usage: "name of the credential",
		},
	}
	FlRegion = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "region",
			Usage: "name of the region",
		},
	}
	FlAvailabilityZoneOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "availability-zone",
			Usage: "name of the availability zone",
		},
	}
	FlImageCatalog = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "imagecatalog",
			Usage: "name of the imagecatalog",
		},
	}
	FlKey = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "key",
			Usage: "key of the tag",
		},
	}
	FlValue = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "value",
			Usage: "value of the tag",
		},
	}
	FlProxyHost = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "proxy-host",
			Usage: "hostname or ip of the proxy",
		},
	}
	FlProxyPort = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "proxy-port",
			Usage: "port of the proxy",
		},
	}
	FlProxyProtocol = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "proxy-protocol",
			Usage: "protocol of the proxy (http or https)",
			Value: "http",
		},
	}
	FlProxyUser = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "proxy-user",
			Usage: "user for the proxy if basic auth is required",
		},
	}
	FlProxyPassword = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "proxy-password",
			Usage: "password for the proxy if basic auth is required",
		},
	}
	FlRdsUserName = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "db-username",
			Usage: "username to use for the jdbc connection",
		},
	}
	FlRdsPassword = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "db-password",
			Usage: "password to use for the jdbc connection",
		},
	}
	FlRdsDriverOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "driver",
			Usage: "[DEPRECATED] has no effect",
		},
	}
	FlRdsURL = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "url",
			Usage: "JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>",
		},
	}
	FlRdsDatabaseEngineOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "database-engine",
			Usage: "[DEPRECATED] has no effect",
		},
	}
	FlRdsType = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "type",
			Usage: "type of database, aka the service name that will use the db like HIVE, DRUID, SUPERSET, RANGER, etc.",
		},
	}
	FlRdsValidatedOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "not-validated",
			Usage: "[DEPRECATED] has no effect, use 'cb database test ...' command instead",
		},
	}
	FLMpackURL = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "url",
			Usage: "URL of the mpack",
		},
	}
	FLMpackPurge = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "purge",
			Usage: "purge existing resources specified in purge-list",
		},
	}
	FLMpackPurgeList = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "purge-list",
			Usage: "comma separated list of resources to purge (stack-definitions,service-definitions,mpacks). By default (stack-definitions,mpacks) will be purged",
		},
	}
	FLMpackForce = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "force",
			Usage: "force install management pack",
		},
	}
	FlRdsConnectorJarURLOptional = StringFlag{
		RequiredFlag: OPTIONAL,
		StringFlag: cli.StringFlag{
			Name:  "connector-jar-url",
			Usage: "URL of the jdbc jar file",
		},
	}
	FlWithCustomDomainOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-custom-domain",
			Usage: "adds custom domain configuration to the template",
		},
	}
	FlWithTagsOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-tags",
			Usage: "adds user defined tags configuration to the template",
		},
	}
	FlWithImageOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-image",
			Usage: "adds image-catalog configuration to the template",
		},
	}
	FlWithKerberosManagedOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-kerberos-managed",
			Usage: "adds Cloudbreak managed Kerberos configuration to the template",
		},
	}
	FlWithKerberosExistingMITOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-kerberos-mit",
			Usage: "adds existing MIT Kerberos configuration to the template",
		},
	}
	FlWithKerberosExistingADOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-kerberos-ad",
			Usage: "adds existing Active Directory Kerberos configuration to the template",
		},
	}
	FlWithKerberosCustomOptional = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-kerberos-custom",
			Usage: "adds custom Kerberos configuration to the template",
		},
	}
	FlWithBlueprintValidation = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "with-blueprint-validation",
			Usage: "enable Ambari blueprint validation",
		},
	}
	FlWithSourceCluster = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "source-cluster",
			Usage: "source cluster to use as datalake",
		},
	}
	FlHostGroups = StringFlag{
		RequiredFlag: REQUIRED,
		StringFlag: cli.StringFlag{
			Name:  "host-groups",
			Usage: "comma separated list of hostgroups where the failed nodes will be repaired",
		},
	}
	FlRemoveOnly = BoolFlag{
		RequiredFlag: OPTIONAL,
		BoolFlag: cli.BoolFlag{
			Name:  "remove-only",
			Usage: "the failed nodes will only be removed, otherwise the failed nodes will be removed and new nodes will be started.",
		},
	}
)

type RequiredFlag struct {
	Required bool
}

type StringFlag struct {
	cli.StringFlag
	RequiredFlag
}

type BoolFlag struct {
	cli.BoolFlag
	RequiredFlag
}

type IntFlag struct {
	cli.IntFlag
	RequiredFlag
}

func requiredFlags(flags []cli.Flag) []cli.Flag {
	required := []cli.Flag{}
	for _, flag := range flags {
		if isRequiredVisible(flag) {
			required = append(required, flag)
		}
	}
	return required
}

func optionalFlags(flags []cli.Flag) []cli.Flag {
	required := []cli.Flag{}
	for _, flag := range flags {
		if flag.GetName() != "generate-bash-completion" {
			if !isRequiredVisible(flag) {
				required = append(required, flag)
			}
		}
	}
	return required
}

func checkRequiredFlagsAndArguments(c *cli.Context) {
	checkRequiredFlags(c)
	argumentsNotAllowed(c)
}

func checkRequiredFlags(c *cli.Context) {
	missingFlags := make([]string, 0)
	for _, f := range c.Command.Flags {
		if isRequired(f) && len(c.String(f.GetName())) == 0 {
			missingFlags = append(missingFlags, f.GetName())
		}
	}
	if len(missingFlags) > 0 {
		utils.LogMissingParameterAndExit(c, missingFlags)
	}
}

func argumentsNotAllowed(c *cli.Context) {
	args := c.Args()
	if len(args) > 0 {
		utils.LogErrorAndExit(fmt.Errorf("Argument %q is not allowed for this command", args.Get(0)))
	}
}

func isRequiredVisible(flag cli.Flag) bool {
	if flag.GetName() == "help, h" || flag.GetName() == "generate-bash-completion" {
		return false
	}
	hidden := flagValue(flag).FieldByName("Hidden").Bool()
	required := flagValue(flag).FieldByName("Required").Bool()
	return !hidden && required
}

func isRequired(flag cli.Flag) (required bool) {
	defer func() {
		if r := recover(); r != nil {
			required = false
		}
	}()
	if flag.GetName() == "help, h" {
		return false
	}
	return flagValue(flag).FieldByName("Required").Bool()
}

func flagValue(flag cli.Flag) reflect.Value {
	fv := reflect.ValueOf(flag)
	for fv.Kind() == reflect.Ptr {
		fv = reflect.Indirect(fv)
	}
	return fv
}

type FlagBuilder struct {
	flags []cli.Flag
}

func NewFlagBuilder() *FlagBuilder {
	return &FlagBuilder{flags: make([]cli.Flag, 0)}
}

func (fb *FlagBuilder) AddFlags(flags ...cli.Flag) *FlagBuilder {
	for _, f := range flags {
		fb.flags = append(fb.flags, f)
	}
	return fb
}

func (fb *FlagBuilder) AddAuthenticationFlags() *FlagBuilder {
	for _, f := range []cli.Flag{FlServerOptional, FlUsername, FlPassword, FlProfileOptional, FlAuthTypeOptional} {
		fb.flags = append(fb.flags, f)
	}
	return fb
}

func (fb *FlagBuilder) AddResourceDefaultFlags() *FlagBuilder {
	for _, f := range []cli.Flag{FlName, FlDescriptionOptional, FlPublicOptional} {
		fb.flags = append(fb.flags, f)
	}
	return fb
}

func (fb *FlagBuilder) AddResourceFlagsWithOptionalName() *FlagBuilder {
	for _, f := range []cli.Flag{FlNameOptional, FlDescriptionOptional, FlPublicOptional} {
		fb.flags = append(fb.flags, f)
	}
	return fb
}

func (fb *FlagBuilder) AddOutputFlag() *FlagBuilder {
	fb.flags = append(fb.flags, FlOutputOptional)
	return fb
}

func (fb *FlagBuilder) AddTemplateFlags() *FlagBuilder {
	for _, f := range []cli.Flag{FlWithCustomDomainOptional, FlWithTagsOptional, FlWithImageOptional, FlWithKerberosManagedOptional, FlWithKerberosExistingMITOptional, FlWithKerberosExistingADOptional, FlWithKerberosCustomOptional, FlWithBlueprintValidation} {
		fb.flags = append(fb.flags, f)
	}
	return fb
}

func (fb *FlagBuilder) Build() []cli.Flag {
	return fb.flags
}
