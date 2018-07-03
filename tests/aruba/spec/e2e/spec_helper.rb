require 'aruba/rspec'
require "rspec/json_expectations"
require 'json'
require 'date'
require "common/command_builder.rb"

def html_print(&blk)
  puts "<pre>"
  blk.call
  puts "</pre>"
end