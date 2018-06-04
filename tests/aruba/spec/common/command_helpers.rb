RSpec.shared_context "shared command helpers", :a => :b do
  before { @some_var = :some_value }

  def scale(cb, cluster_name, hostgroup, scale_range, &block)
    skip_if(cb, cluster_name, "AVAILABLE", "Test is skipped because of cluster is not AVAILABLE")       
    before_node_count = get_node_count(cb, cluster_name, hostgroup)
    expect(before_node_count).to be >= 0
     
    result = cb.cluster.scale.name(cluster_name).group_name(hostgroup).desired_node_count(before_node_count + scale_range).build            
    expect(result.exit_status).to eql 0

    result = block.call

    after_node_count = get_node_count(cb, cluster_name, hostgroup)
    expect(after_node_count).to eq(before_node_count + scale_range) 
  end

  def bp_create_describe_delete(cb, blueprint_name, &block)
    result = block.call
    expect(result.exit_status).to eql 0

    result = cb.blueprint.describe.name(blueprint_name).build 
    expect(result.exit_status).to eql 0    

    result = cb.blueprint.delete.name(blueprint_name).build 
    expect(result.exit_status).to eql 0        
  end 

   def recipe_create_describe_list_delete(cb, recipe_name, mock= false, &block)
    result = block.call
    expect(result.exit_status).to eql 0

    result = cb.recipe.describe.name(recipe_name).build 
    expect(result.exit_status).to eql 0

    if mock 
      result = cb.recipe.list.build 
      expect(result.exit_status).to eql 0   
    else   
      result = recipe_list_with_check(recipe_name)
      expect(result).to be_truthy
    end             
    result = cb.recipe.delete.name(recipe_name).build 
    expect(result.exit_status).to eql 0        
  end

  def credential_create_describe_list_delete(cb, cred_name, &block)
    result = block.call
    expect(result.exit_status).to eql 0

    result = cb.credential.describe.name(cred_name).build 
    expect(result.exit_status).to eql 0

    result = credential_list_with_check(cred_name)
    expect(result).to be_truthy            

    result = cb.credential.delete.name(cred_name).build 
    expect(result.exit_status).to eql 0        
  end

  def recipe_list_with_check(recipe_name) 
    got_recipe = false

    result = cb.recipe.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    json = JSON.parse(result.stdout)

    json.each do |s|
       if s["Name"] == recipe_name
        got_recipe = true
        return true
       end
          
      expect(s).to include_json(
        Name: /.*/,
        Description: /.*/,  
        ExecutionType: /.*/     
    )
    end 
    return false  
  end

  def recipe_exist(json, recipe_name)
    got_recipe = false
    json.each do |s|
       if s["Name"] == recipe_name
        got_recipe = true
        return true
       end
    end
    return false   
  end  

  def blueprint_list_with_check(bp_name) 
    got_bp = false

    result = cb.blueprint.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    json = JSON.parse(result.stdout)

    json.each do |s|
       if s["Name"] == bp_name
        got_bp = true
        return true
       end
    end 
    return false  
  end

  def credential_list_with_check(cred_name) 
    got_cred = false

    result = cb.credential.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    json = JSON.parse(result.stdout)

    json.each do |s|
       if s["Name"] == cred_name
        got_cred = true
        return true
       end
    end 
    return false  
  end

  def imagecat_list_with_check(ic_name) 
    got_ic = false

    result = cb.imagecatalog.list.build 
    expect(result.exit_status).to eql 0 
    
    expect(result.stdout.empty?).to be_falsy
    json = JSON.parse(result.stdout)

    json.each do |s|
       if s["Name"] == ic_name
        got_cred = true
        return true
       end
    end 
    return false  
  end          

  def get_region(result)
    expect(result.empty?).to be_falsy
    json = JSON.parse(result)
    json.each  do |a|
      return a["Name"]
    end
  end

  def create_p12(file_name)
    result = run("touch " + file_name)
    result.stop
    expect(result.exit_status).to eql 0       
  end

  def create_valid_json(file_name, json)
    f = File.open(File.dirname(__FILE__) + "/../../tmp/aruba/" + file_name,"w")
    f.write(json)
    f.close
  end 

  def get_cluster_info(cb, cluster_name)
    json = get_cluster_json(cb, cluster_name)
    expect(json.empty?).to be_falsy
      html_print do 
        puts "\nCLUSTER INFORMATION:"
        puts "Cluster Name:    " + json["name"].to_s
        puts "Stack Id:        " + json["id"].to_s + "\n"     
      end   
  end  

  let(:shared_let) { {'arbitrary' => 'object'} } 
  subject do
    'this is the subject'
  end
end