require_relative "../common/mock_vars"
require_relative "../common/command_helpers"
require_relative "spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "cb "
  return cb
end

RSpec.describe 'Blueprint test cases', :type => :aruba do
  include_context "shared command helpers"    
  include_context "mock shared vars"

  it "Blueprint - Create from url" do 
    with_environment 'DEBUG' => '1' do
      result = cb.blueprint.create.from_url.name("testbp").url(@blueprint_url).build(false)
      expect(result.exit_status).to eql 0
      expect(result.stderr).to include("blueprint created")    
    end
  end    

  it "Blueprint - Create - Url doesn't exist" do
    result = cb.blueprint.create.from_url.name("temp-bp").url("https://something123456789.com").build(false)
    expect(result.exit_status).to eql 1
    expect(result.stderr).to include("error") 
  end

  it "Blueprint - Create - Invalid url with no protocol " do
    result = cb.blueprint.create.from_url.name("temp-bp").url("something123456789.com").build(false)
    expect(result.exit_status).to eql 1
    expect(result.stderr).to include("error") 
  end

  it "Blueprint - Create from file" do
    with_environment 'DEBUG' => '1' do
      result = cb.blueprint.create.from_file.name(@blueprint_name_file).file(@blueprint_file).build(false)
      expect(result.exit_status).to eql 0
      expect(result.stderr).to include("blueprint created")
    end
  end

  it "Blueprint - Describe a default blueprint" do
    result = cb.blueprint.describe.name(@default_blueprint_name).build(false)
    expect(result.exit_status).to eql 0
    expect(result.stdout.empty?).to be_falsy
    expect(JSON.parse(result.stdout)).to include_json(
      Name: /.*/,
      Description: /.*/,
      HDPVersion: /.*/,
      HostgroupCount: /.*/,
      Tags: /.*/
    )
  end

  it "Blueprint - List - All existing" do
    result = cb.blueprint.list.build(false)
    expect(result.exit_status).to eql 0
    expect(result.stdout.empty?).to be_falsy
    JSON.parse(result.stdout).each do |s|    
      expect(s).to include_json(
        Name: /.*/,
        Description: /.*/,  
        HDPVersion: /.*/,
        HostgroupCount: /.*/,
        Tags: /.*/     
    )
    end       
  end              
end  