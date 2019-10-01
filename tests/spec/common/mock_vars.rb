RSpec.shared_context "mock shared vars", :a => :b do
  before (:all) {
    @env_name = "mockenv"

    @aws_credential_name = "amazon"

    @aws_role_arn = "arn:aws:iam::1234567890:role/auto-test"
    @aws_access_key = "ABCDEFGHIJKLMNO123BC"
    @aws_secret_key = "+ABcdEFGHIjKLmNo+ABc0dEF1G23HIjKLmNopqrs"

    @mock_password = "mockpassword"

    @mock_endpoint_setup = "/cb/api/setup"
    @mock_endpoint_reset = "/cb/api/reset"
    @mock_password = "mockpassword"

    @environment_file = "../../templates/aws-environment-skeleton.json"
    @environment_crn = "crn:cdp:environments:us-west-1:hortonworks:environment:a123bcde-fgh4-5i6j-7k8l-m90nop124qrs"

    @dbserver_crn = "crn:cdp:redbeams:us-west-1:hortonworks:databaseServer:a123bcde-fgh4-5i6j-7k8l-m90nop124qrs"
    @dbserver_name = "dbstck-mock-amazon-12345a6b7cd8e9012"
    @dbserver_create_file = "../../templates/dbserver-create-skeleton.json"
    @dbserver_register_file = "../../templates/dbserver-register-skeleton.json"
    @db_crn = "crn:cdp:redbeams:us-west-1:hortonworks:database:12345ab6-7c89-012d-e343-fgh567iij890"
    @db_name = "mock-db"
    @db_create_file = "../../templates/db-create-skeleton.json"
    @db_register_file = "../../templates/db-register-skeleton.json"
  }
end
