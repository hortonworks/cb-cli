require 'rest-client'
require 'json'
require_relative "../integration/spec_helper"

class TraceResponseBuilder
    @@default_workspace = 1

    @@cloudbreak_base = "cb"
    @@api_base = "#{@@cloudbreak_base}/api"
    @@blueprint_base = "#{@@api_base}/v3/#{@@default_workspace}/blueprints"
    @@create_blueprint_endpoint = "#{@@blueprint_base}"
    @@cluster_base = "#{@@api_base}/v3/#{@@default_workspace}/stack"
    @@create_cluster_endpoint = "#{@@cluster_base}"
    @@create_workspace_endpoint = "#{@@api_base}/v3/workspaces"
    @@get_workspace_endpoint = "#{@@api_base}/v3/workspaces"
    @@get_users_endpoint = "#{@@api_base}/v1/users"

    def self.createWorkspaceRequestFactory(requestBody)
        return {
            :calledEndpoint => @@create_workspace_endpoint,
            :sentValue => requestBody
        }
    end

    def self.getWorkspacesResponseFactory(responseBody)
        return {
            :calledEndpoint => @@get_workspace_endpoint,
            :receivedValue => responseBody
        }
    end

    def self.getWorkspaceByNameResponseFactory(responseBody)
        return {
            :calledEndpoint => "#{@@get_workspace_endpoint}/name/mock@hortonworks.com",
            :receivedValue => responseBody
        }
    end

    def self.deleteWorkspaceByNameResponseFactory(responseBody)
        return {
            :calledEndpoint => "#{@@get_workspace_endpoint}/name/mock@hortonworks.com",
            :receivedValue => responseBody
        }
    end

    def self.getAllUsersResponseFactory(responseBody)
        return {
            :calledEndpoint => @@get_users_endpoint,
            :receivedValue => responseBody
        }
    end
end
