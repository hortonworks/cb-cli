var responses={};

var acc = require('./accountpreferences/accountpreferences.json');
var platforms = require('./accountpreferences/platforms.json');
var defaultblueprint = require('./blueprints/default-blueprint.json');
var qablueprint = require('./blueprints/qa-blueprints.json');
var disktypes = require('./connectors/disktypes.json');
var gateways = require('./connectors/gateways.json');
var ippools = require('./connectors/ippools.json');
var connetworks = require('./connectors/networks.json');
var recommendations = require('./connectors/recommendations.json');
var regions = require('./connectors/regions.json');
var sshkeys = require('./connectors/sshkeys.json');
var securitygroups = require('./connectors/securitygroups.json');
var credentialaws = require('./credentials/aws.json');
var credentialazure = require('./credentials/azure.json');
var credentialgcp = require('./credentials/gcp.json');
var credentialopenstack = require('./credentials/openstack.json');
var cbdefimagec = require('./imagecatalogs/cloudbreak-default.json');
var defimagec = require('./imagecatalogs/default-imagecatalog.json');
var qaimages = require('./imagecatalogs/qa-images.json');
var ldap = require('./ldapconfig/default-ldap.json');
var networks = require('./networks/networks.json');
var proxy = require('./proxyconfig/default-proxy.json');
var rds = require('./rdsconfig/default-rds.json');
var recipes = require('./recipes/recipes.json');
var responses = require('./responses.js');
var secgr = require('./securitygroups/securitygroups.json');
var secrules = require('./securityrules/securityrules.json');
var aws = require('./stacks/aws.json');
var azure = require('./stacks/azure.json');
var gcp = require('./stacks/gcp.json');
var openstack = require('./stacks/openstack.json');
var datalake = require('./stacks/datalake.json');
var templates = require('./templates/qa-templates.json');
var profile = require('./users/default-profile.json');
var matrix = require('./utilsmatrix.json');
var mpacks = require('./mpacks/mpacks.json');
var organizations = require('./organizations/organizations.json');

const OK = 200;

var responseObject = function(payload, code) {
	return { responses : [ { response: payload , statusCode: code } ] };
}

var addResponseObject = function(responses, payload, code, conditionText) {
	var newElement = { response: payload, statusCode: code, condition: conditionText };
	responses.responses.push(newElement);
	return responses;
}

var responseObjectWithCondition = function(payload, code, conditionText) {
	return { responses : [ { response: payload , statusCode: code , condition: conditionText} ] };
}

var stackResponses = responseObjectWithCondition({"message":"Stack 'az404' not found"}, 404, "return params['name'].value === 'az404';");
stackResponses = addResponseObject(stackResponses, openstack, OK, "return (params['name'].value !== 'az404' && params['name'].value !== 'dl-ok')");
stackResponses = addResponseObject(stackResponses, datalake, OK, "return params['name'].value === 'dl-ok';");

var stackOperationResponses = responseObjectWithCondition({"message":"Stack 'azstatus' not found"}, 404, "return params['name'].value === 'azstatus';");
stackOperationResponses = addResponseObject(stackOperationResponses,null, OK , "return params['name'].value !== 'azstatus';");

var stackReinstallResponses = responseObjectWithCondition({"message":"Stack 'aaaaa' not found"}, 404, "return params['name'].value === 'aaaaa';");
stackReinstallResponses = addResponseObject(stackReinstallResponses,null, OK , "return params['name'].value !== 'aaaaa';");

responses.getCloudbreakInfo= responseObject({ "app": { "name":"cloudbreak", "version":"MOCK" } }, OK);
responses.getCloudbreakHealth= responseObject({ "status":"UP" }, OK);
responses.getAccountPreferencesEndpoint= responseObject(acc, OK);
responses.isPlatformSelectionDisabled= responseObject(platforms, OK);
responses.platformEnablement= responseObject(platforms, OK);
responses.listBlueprintsByOrganization = responseObject(qablueprint, OK);
responses.getBlueprintInOrganization = responseObject(defaultblueprint, OK);
responses.getPrivatesBlueprint= responseObject(qablueprint, OK);
responses.getPrivateBlueprint= responseObject(defaultblueprint, OK);
responses.getBlueprint= responseObject(defaultblueprint, OK);
responses.createBlueprintInOrganization= responseObject(defaultblueprint, OK);
responses.postPrivateBlueprint= responseObject(defaultblueprint, OK);
responses.getBlueprintRequestFromId= responseObject(defaultblueprint, OK);
responses.getDisktypes= responseObject(disktypes, OK);
responses.getGatewaysCredentialId= responseObject(gateways, OK);
responses.getIpPoolsCredentialId= responseObject(ippools, OK);
responses.getPlatformNetworks= responseObject(connetworks, OK);
responses.getPublicsNetwork= responseObject(networks, OK);
responses.createRecommendation= responseObject(recommendations, OK);
responses.getRegions= responseObject(regions, OK);
responses.getPlatformSShKeys= responseObject(sshkeys, OK);
responses.getCredentialInOrganization = responseObject(credentialopenstack, OK);
responses.createCredentialInOrganization = responseObject(credentialopenstack, OK);
responses.listCredentialsByOrganization = responseObject([credentialopenstack, credentialgcp, credentialazure, credentialaws], OK);
responses.getPublicCredential= responseObject(credentialopenstack, OK);
responses.getPrivatesCredential= responseObject([credentialopenstack,credentialgcp,credentialazure,credentialaws], OK);
responses.getPrivateCredential= responseObject(credentialopenstack, OK);
responses.listImageCatalogsByOrganization = responseObject(defimagec, OK);
responses.getImageCatalogInOrganization = responseObject(defimagec, OK);
responses.createImageCatalogInOrganization = responseObject(defimagec, OK);
responses.getPublicImageCatalogsByName= responseObject(defimagec, OK);
responses.getImagesByProviderAndCustomImageCatalogInOrganization= responseObject(qaimages, OK);
responses.putSetDefaultImageCatalogByName= responseObject(defimagec, OK);
responses.postPrivateImageCatalog= responseObject(defimagec, OK);
responses.getImageCatalogRequestFromName= responseObject(defimagec, OK);
responses.listLdapsByOrganization = responseObject(ldap, OK);
responses.getLdapConfigInOrganization = responseObject(ldap, OK);
responses.createLdapConfigsInOrganization = responseObject(ldap, OK);
responses.postLdapConnectionTestInOrganization = responseObject({ "connectionResult":"Failed to connect to LDAP server: hwxad-1a2bcd3e45678f90.elb.eu-west-1.amazonaws.com:123" }, OK);
responses.getPrivatesLdap= responseObject(ldap, OK);
responses.postPrivateLdap= responseObject(ldap, OK);
responses.getPrivateLdap= responseObject(ldap, OK);
responses.getLdap= responseObject(ldap, OK);
responses.listProxyconfigsByOrganization = responseObject(proxy, OK);
responses.getProxyconfigInOrganization = responseObject(proxy, OK);
responses.createProxyconfigInOrganization = responseObject(proxy, OK);
responses.getPrivatesProxyConfig= responseObject(proxy, OK);
responses.postPrivateProxyConfig= responseObject(proxy, OK);
responses.getPrivateProxyConfig= responseObject(proxy, OK);
responses.getProxyConfig= responseObject(proxy, OK);
responses.listRdsConfigsByOrganization = responseObject(rds, OK);
responses.createRdsConfigInOrganization = responseObject(rds, OK);
responses.getRdsConfigInOrganization = responseObject(rds, OK);
responses.testRdsConnectionInOrganization = responseObject({ "connectionResult":"Failed to connect to RDS: The connection attempt failed." }, OK);
responses.getPrivatesRds= responseObject(rds, OK);
responses.postPrivateRds= responseObject(rds, OK);
responses.getPrivateRds= responseObject(rds, OK);
responses.listRecipesByOrganization= responseObject(recipes, OK);
responses.getPublicRecipe= responseObject(recipes[0], OK);
responses.postPublicRecipe= responseObject(recipes[0], OK);
responses.postPrivateRecipe= responseObject(recipes[0], OK);
responses.getRecipeRequestFromName = responseObject(recipes, OK);
responses.getPublicsSecurityGroup = responseObject(secgr, OK);
responses.getRecipeInOrganization = responseObject(recipes[0], OK);
responses.getDefaultSecurityRules= responseObject(secrules, OK);
responses.getPublicsStack= responseObject([openstack,aws,azure,gcp], OK);
responses.getPublicStack= stackResponses;
responses.getPublicsTemplate= responseObject(templates, OK);
responses.getUserProfile= responseObject(profile, OK);
responses.getRegionsByCredentialId= responseObject(regions, OK);
responses.listStacksByOrganization = responseObject([openstack,aws,azure,gcp], OK);
responses.createStackInOrganization = responseObject(openstack, OK);
responses.getStackInOrganization = responseObject(openstack, OK);
responses.getPublicStackV2= responseObject(stackResponses, OK);
responses.postPrivateStackV2= responseObject(openstack, OK);
responses.getPrivateStackV2= responseObject(openstack, OK);
responses.getStackV2= responseObject(openstack, OK);
responses.getPlatformSecurityGroups= responseObject(securitygroups, OK);
responses.createManagementPackInOrganization = responseObject(mpacks, OK);
responses.postPrivateManagementPack = responseObject(mpacks, OK);
responses.getOrganizations = responseObject(organizations, OK);

responses.postRepositoryConfigsValidation= responseObject({
  "utilsBaseURL" : true,
  "ambariGpgKeyUrl" : true,
  "mpackUrl" : true,
  "stackBaseURL" : true,
  "versionDefinitionFileUrl" : true,
  "ambariBaseUrl" : true
}, OK);
responses.getStackMatrixUtil= responseObject(matrix , OK);
responses.getDefaultSmartSenseSubscriptionInOrganization= responseObject({
  "owner" : "owner",
  "publicInAccount" : false,
  "id" : 1,
  "autoGenerated" : false,
  "subscriptionId" : "subscriptionId",
  "account" : "account"
}, OK);
responses.getTagSpecifications = responseObject({ "specifications" : { "key" : { "key" : "{}" } } }, OK);
responses.listManagementPacksByOrganization = responseObject([ {
  "public" : false,
  "name" : "name",
  "description" : "description",
  "purge" : false,
  "force" : false,
  "id" : 0,
  "mpackUrl" : "mpackUrl",
  "purgeList" : [ "purgeList", "purgeList" ]
}, {
  "public" : false,
  "name" : "name",
  "description" : "description",
  "purge" : false,
  "force" : false,
  "id" : 0,
  "mpackUrl" : "mpackUrl",
  "purgeList" : [ "purgeList", "purgeList" ]
} ],OK);
responses.listFlexSubscriptionsByOrganization = responseObject([ {
  "owner" : "owner",
  "publicInAccount" : false,
  "smartSenseSubscriptionId" : 0,
  "usedForController" : false,
  "name" : "name",
  "smartSenseSubscription" : {
    "owner" : "owner",
    "publicInAccount" : false,
    "id" : 1,
    "autoGenerated" : false,
    "subscriptionId" : "subscriptionId",
    "account" : "account"
  },
  "id" : 6,
  "subscriptionId" : "subscriptionId",
  "account" : "account",
  "usedAsDefault" : false
}, {
  "owner" : "owner",
  "publicInAccount" : false,
  "smartSenseSubscriptionId" : 0,
  "usedForController" : false,
  "name" : "name",
  "smartSenseSubscription" : {
    "owner" : "owner",
    "publicInAccount" : false,
    "id" : 1,
    "autoGenerated" : false,
    "subscriptionId" : "subscriptionId",
    "account" : "account"
  },
  "id" : 6,
  "subscriptionId" : "subscriptionId",
  "account" : "account",
  "usedAsDefault" : false
} ], OK);

responses.putstartStackV2 = stackOperationResponses;
responses.putscalingStackV2 = stackOperationResponses;
responses.putstopStackV2 = stackOperationResponses;
responses.putrepairStackV2 = stackOperationResponses;
responses.putsyncStackV2 = stackOperationResponses;
responses.putreinstallStackV2 = stackReinstallResponses;

responses.checkClientVersion = responseObject({ "versionCheckOk" : true, "message" : "message" }, OK);
 
module.exports= {
  responses:responses
}
