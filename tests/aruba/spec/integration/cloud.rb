require "common/e2e_vars.rb"
require "common/helpers.rb"
require "common/command_helpers.rb"
require "e2e/spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "cb "
  return cb
end

RSpec.describe 'Cloud test cases', :type => :aruba do
  include_context "shared helpers"
  include_context "shared command helpers"    
  include_context "shared vars"

  it "Cloud - Availability zones list" do 
      result = cb.cloud.availability_zones.credential(@os_credential_name + "-cloud").region("region").build()
      expect(result.exit_status).to eql 0
      JSON.parse(result.stdout).each do |s| 
        expect(s).to include_json(
          Name: /.*/  
        ) 
      end
  end

  it "Cloud - Availability zones list - No Credential/Region" do 
      result = cb.cloud.availability_zones.credential("").region("").build
      expect(result.exit_status).to eql 1
  end  

   it "Cloud - Instances list" do 
      result = cb.cloud.instances.credential(@os_credential_name + "-cloud").region("region").build
      expect(result.exit_status).to eql 0
  end 

   it "Cloud - Instances list - No Credential/Region" do 
      result = cb.cloud.instances.credential("").region("").build
      expect(result.exit_status).to eql 1
  end 

  it "Cloud - Volumes list AWS" do 
    result = cb.cloud.volumes.aws.build
    expect(result.exit_status).to eql 0
    JSON.parse(result.stdout).each do |s| 
      expect(s).to include_json(
        Name: /.*/,
        Description: /.*/  
      ) 
    end
  end  

  it "Cloud - Volumes list Azure" do 
    result = cb.cloud.volumes.azure.build
    expect(result.exit_status).to eql 0
    JSON.parse(result.stdout).each do |s| 
      expect(s).to include_json(
        Name: /.*/,
        Description: /.*/  
      ) 
    end
  end              

  it "Cloud - Volumes list GCP" do 
    result = cb.cloud.volumes.gcp.build
    expect(result.exit_status).to eql 0
    JSON.parse(result.stdout).each do |s| 
      expect(s).to include_json(
        Name: /.*/,
        Description: /.*/  
      ) 
    end
  end   
end  