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
end