require "common/e2e_vars.rb"
require "common/helpers.rb"
require "common/command_helpers.rb"
require "e2e/spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "cb "
  return cb
end

RSpec.describe 'Blueprint test cases', :type => :aruba do
  include_context "shared helpers"
  include_context "shared command helpers"    
  include_context "shared vars"

  it "Blueprint - Create from url - Describe - Delete " do 
    bp_create_describe_delete(cb, @blueprint_name_url) do
      cb.blueprint.create.from_url.name(@blueprint_name_url).url(@blueprint_url).build
    end 
  end    

  it "Blueprint - Create - Url doesn't exist" do
    result = cb.blueprint.create.from_url.name("temp-bp").url("https://something123456789.com").build  
    expect(result.exit_status).to eql 1
    expect(result.stderr).to include("error") 
  end

  it "Blueprint - Create - Invalid url with no protocol " do
    result = cb.blueprint.create.from_url.name("temp-bp").url("something123456789.com").build  
    expect(result.exit_status).to eql 1
    expect(result.stderr).to include("error") 
  end

  it "Blueprint - Create from file - Describe - Delete " do 
    bp_create_describe_delete(cb, @blueprint_name_file) do
      cb.blueprint.create.from_file.name(@blueprint_name_file).file(@blueprint_file).build  
    end 
  end 

  it "Blueprint - Describe a default blueprint" do
    result = cb.blueprint.describe.name(@default_blueprint_name).build 
    expect(result.exit_status).to eql 0
    expect(JSON.parse(result.stdout)).to include_json(
      Name: /.*/,
      Description: /.*/,  
      HDPVersion: /.*/,
      HostgroupCount: /.*/,
      Tags: /.*/     
    )       
  end

  it "Blueprint - List - All existing" do
    result = cb.blueprint.list.build 
    expect(result.exit_status).to eql 0

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