RSpec.shared_context "e2e shared vars", :a => :b do
  before (:all) {
  @os_cluster_name = "cli-os-cluster"
  @os_credential_name = "cli-os-cluster"
  @cli_input_json = "../../templates/kilo-openstack-template.json"
  @aws_credential_name = "cli-aws-cred"

  @environment_name = "cli-environment"
  @environment_location = "Frankfurt"
  @environment_regions = "RegionOne"

  @ambari_user = "admin"
  @ambari_password = 'Admin123!@#\"'

  @recipe_types = Array["pre-cluster-manager-start", "pre-termination", "post-cluster-manager-start", "post-cluster-install"]
  @recipe_name = "cli-recipe"
  @recipe_file = "../../recipes/echo.sh"

  @default_blueprint_name = "'HDP 3.1 - Data Science: Apache Spark 2, Apache Zeppelin'"
  @blueprint_name_url = "cli-bp-url"
  @blueprint_name_file = "cli-bp-file"
  @blueprint_url = ENV['CLUSTER_DEFINITION_URL']
  @blueprint_file = "../../blueprints/test.bp"

  @image_catalog_name = "cli-cat"
  @image_catalog_name_default = "cloudbreak-default"
  @image_catalog_url = "https://s3-us-west-1.amazonaws.com/cloudbreak-imagecatalog/v2-dev-cb-image-catalog.json"

  @db_name = "cli-db-postgr"

  @ldap_name = "cli-ldap"
  @ldap_domain = "ad.hwx.com"
  @ldap_bind_dn = "CN=Administrator,CN=Users,DC=ad,DC=hwx,DC=com"
  @ldap_user_search_base = "CN=Users,DC=ad,DC=hwx,DC=com"
  @ldap_user_dn_pattern= "CN={0},CN=Users,DC=ad,DC=hwx,DC=com"
  @ldap_user_name_attribute = "sAMAccountName"
  @ldap_user_object_class = "person"
  @ldap_group_member_attribute = "member"
  @ldap_group_name_attribute = "cn"
  @ldap_group_object_class = "group"
  @ldap_group_search_base = "CN=Users,DC=ad,DC=hwx,DC=com"

  @proxy_name = "cli-proxy"
  @proxy_server = "1.2.3.4"
  @proxy_port = "123"
  @proxy_user = "fakeuser"
  @proxy_password = "fakepwd"

  @mpack_name = "cli-mpack"
  @mpack_url = "http://public-repo-1.hortonworks.com/HDF/centos6/3.x/updates/3.1.1.0/tars/hdf_ambari_mp/hdf-ambari-mpack-3.1.1.0-35.tar.gz"
  }
end
