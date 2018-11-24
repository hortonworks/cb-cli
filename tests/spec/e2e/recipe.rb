require_relative "../common/e2e_vars"
require_relative "../common/command_helpers"
require_relative "spec_helper"

define_method(:cb) do
  cb = CommandBuilder.new
  CommandBuilder.cmd = "dp "
  return cb
end

RSpec.describe 'Recipe test cases', :type => :aruba do
  include_context "shared command helpers"    
  include_context "e2e shared vars"

  before(:all) do
    result = cb.recipe.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    recipe_list_json = JSON.parse(result.stdout)

    @recipe_types.each  do |type|
      r_name = "cli-" + type + "-u"
      recipe_exist = json_has_name(recipe_list_json, r_name)

      if (recipe_exist[0] == true)
        result = cb.recipe.delete.name(r_name).build
        expect(result.exit_status).to eql 0 
      end
    end  
  end

  before(:all) do
    result = cb.recipe.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    recipe_list_json = JSON.parse(result.stdout)

    @recipe_types.each  do |type|
      r_name = "cli-" + type + "-f"
      recipe_exist = json_has_name(recipe_list_json, r_name)

      if (recipe_exist[0] == true)
        result = cb.recipe.delete.name(r_name).build
        expect(result.exit_status).to eql 0 
      end
    end  
  end

  it "Recipe - Create from url - Describe - List - Delete - All recipe types" do 
    @recipe_types.each  do |type|
      r_name = "cli-" + type + "-u"
      recipe_create_describe_list_delete(cb, r_name) do
        cb.recipe.create.from_url.name(r_name).execution_type(type).url(@recipe_url).build 
      end
    end 
  end

   it "Recipe - Create from file - Describe - List - Delete - All recipe types" do 
    @recipe_types.each  do |type|
      r_name = "cli-" + type + "-f"
      recipe_create_describe_list_delete(cb, r_name) do
        cb.recipe.create.from_file.name(r_name).execution_type(type).file(@recipe_file).build 
      end
    end 
  end     
end  