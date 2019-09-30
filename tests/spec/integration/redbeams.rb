require_relative "../common/mock_vars"
require_relative "../common/command_helpers"
require_relative "../common/mock_helper"
require_relative "../common/trace_response_builder"
require_relative "../common/response_helpers"
require_relative "spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "dp "
  return cb
end

RSpec.describe 'Redbeams test cases', :type => :aruba, :feature => "Redbeams", :severity => :critical do
  include_context "shared command helpers"
  include_context "mock shared vars"

  before(:each) do
      MockHelper.resetMock("beams")
  end

  it "Redbeams - Databases - List", :story => "Redbeams", :severity => :normal, :testId => 1 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/dbs.json")
      expectedEndpointResponse = TraceResponseBuilder.listDatabasesResponseFactory(responseHash)
      MockHelper.setupResponse("beams", "listDatabases", responseHash)

      result = cb.redbeams.db.list.env_crn(@environment_crn).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - List", :story => "Redbeams", :severity => :normal, :testId => 2 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/dbservers.json")
      expectedEndpointResponse = TraceResponseBuilder.listDatabaseServersResponseFactory(responseHash)
      MockHelper.setupResponse("beams", "listDatabaseServers", responseHash)

      result = cb.redbeams.dbserver.list.env_crn(@environment_crn).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Describe by CRN", :story => "Redbeams", :severity => :normal, :testId => 3 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/dbserver.json")
      expectedEndpointResponse = TraceResponseBuilder.getDatabaseServerByCrnResponseFactory(responseHash, @dbserver_crn)
      MockHelper.setupResponse("beams", "getDatabaseServerByCrn", responseHash)

      result = cb.redbeams.dbserver.describe.crn(@dbserver_crn).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Describe by Name", :story => "Redbeams", :severity => :normal, :testId => 4 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/dbserver.json")
      expectedEndpointResponse = TraceResponseBuilder.getDatabaseServerByNameResponseFactory(responseHash, @dbserver_name)
      MockHelper.setupResponse("beams", "getDatabaseServerByName", responseHash)

      result = cb.redbeams.dbserver.describe.env_crn(@environment_crn).name(@dbserver_name).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Create new DB Server", :story => "Redbeams", :severity => :critical, :testId => 5 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/create-redbeams-response.json")
      requestHash = MockHelper.getResponseHash("../../../requests/redbeams/create-redbeams-request.json")
      expectedEndpointRequest = TraceResponseBuilder.createDatabaseServerRequestFactory(requestHash)
      MockHelper.setupResponse("beams", "createDatabaseServer", responseHash)

      result = cb.redbeams.dbserver.create.database_server_creation_file(@dbserver_create_file).build(true)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getRequestDiff("beams", expectedEndpointRequest)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Release DB Server", :story => "Redbeams", :severity => :normal, :testId => 6 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/release-redbeams-response.json")
      expectedEndpointResponse = TraceResponseBuilder.releaseManagedDatabaseServerResponseFactory(responseHash, @dbserver_crn)
      MockHelper.setupResponse("beams", "releaseManagedDatabaseServer", responseHash)

      result = cb.redbeams.dbserver.release.crn(@dbserver_crn).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Register DB Server", :story => "Redbeams", :severity => :normal, :testId => 7 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/register-redbeams-response.json")
      requestHash = MockHelper.getResponseHash("../../../requests/redbeams/register-redbeams-request.json")
      expectedEndpointRequest = TraceResponseBuilder.registerDatabaseServerRequestFactory(requestHash)
      MockHelper.setupResponse("beams", "registerDatabaseServer", responseHash)

      result = cb.redbeams.dbserver.register.database_server_registration_file(@dbserver_register_file).build(true)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getRequestDiff("beams", expectedEndpointRequest)).to be_truthy
    end
  end

  it "Redbeams - Database Servers - Delete DB Server by CRN", :story => "Redbeams", :severity => :normal, :testId => 8 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/redbeams/delete-redbeams-response.json")
      expectedEndpointResponse = TraceResponseBuilder.deleteDatabaseServerByCrnResponseFactory(responseHash, @dbserver_crn)
      MockHelper.setupResponse("beams", "deleteDatabaseServerByCrn", responseHash)

      result = cb.redbeams.dbserver.delete.crn(@dbserver_crn).build(false)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end
end