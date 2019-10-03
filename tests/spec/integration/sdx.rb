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

RSpec.describe 'SDX test cases', :type => :aruba, :feature => "SDX", :severity => :critical do
  include_context "shared command helpers"    
  include_context "mock shared vars"   

  before(:each) do
      MockHelper.resetMock("dl")
  end

  it "SDX - List", :story => "SDX", :severity => :normal, :testId => 1 do
    with_environment 'DEBUG' => '1' do
      responseHash = MockHelper.getResponseHash("../../../responses/sdx/sdxes.json")
      expectedEndpointResponse = TraceResponseBuilder.listSdxResponseFactory(responseHash)
      MockHelper.setupResponse("dl", "listSdx", responseHash)

      result = cb.sdx.list.build(true)
      resultHash = MockHelper.getResultHash(result.output)

      expect(result.exit_status).to eql 0
      expect(result.stderr.to_s.downcase).not_to include("error")
      expect(MockHelper.getResponseDiff(expectedEndpointResponse, resultHash)).to be_truthy
    end
  end
end