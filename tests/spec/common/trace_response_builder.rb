require 'rest-client'
require 'json'
require_relative "../integration/spec_helper"

class TraceResponseBuilder
    @@default_workspace = 1

    @@cb_base = "cb"
    @@dl_base = "dl"
    @@env_base = "environmentservice"
    @@beams_base = "redbeams"
    @@cb_api_base = "#{@@cb_base}/api"
    @@dl_api_base = "#{@@dl_base}/api"
    @@env_api_base = "#{@@env_base}/api"
    @@beams_api_base = "#{@@beams_base}/api"
    @@blueprint_base = "#{@@cb_api_base}/v4/#{@@default_workspace}/blueprints"
    @@create_blueprint_endpoint = "#{@@blueprint_base}"
    @@cluster_base = "#{@@cb_api_base}/v4/#{@@default_workspace}/stack"
    @@create_cluster_endpoint = "#{@@cluster_base}"
    @@create_workspace_endpoint = "#{@@cb_api_base}/v4/workspaces"
    @@get_workspace_endpoint = "#{@@cb_api_base}/v4/workspaces"
    @@get_users_endpoint = "#{@@cb_api_base}/v4/users"
    @@base_credential_endpoint = "#{@@env_api_base}/v1/credentials"
    @@prerequisites_endpoint = "#{@@base_credential_endpoint}/prerequisites"
    @@base_environment_endpoint = "#{@@env_api_base}/v1/env"
    @@base_beams_db_endpoint = "#{@@beams_api_base}/v4/databases"
    @@base_beams_dbserver_endpoint = "#{@@beams_api_base}/v4/databaseservers"

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

    def self.listCredentialsResponseFactory(responseBody)
        return {
            :calledEndpoint => @@base_credential_endpoint,
            :receivedValue => responseBody
        }
    end

    def self.getCredentialByNameV1ResponseFactory(responseBody, name)
        return {
            :calledEndpoint => "#{@@base_credential_endpoint}/name/#{name}",
            :receivedValue => responseBody
        }
    end

    def self.createCredentialRequestFactory(requestBody)
        return {
            :calledEndpoint => @@base_credential_endpoint,
            :sentValue => requestBody
        }
    end

    def self.modifyCredentialRequestFactory(requestBody)
        return {
            :calledEndpoint => @@base_credential_endpoint,
            :sentValue => requestBody
        }
    end

    def self.getPrerequisitesByCloudPlatformV1ResponseFactory(cloudPlatform, responseBody)
        return {
            :calledEndpoint => [@@prerequisites_endpoint, cloudPlatform].join("/"),
            :receivedValue => responseBody
        }
    end

    def self.listEnvironmentsResponseFactory(responseBody)
        return {
            :calledEndpoint => @@base_environment_endpoint,
            :receivedValue => responseBody
        }
    end

    def self.createEnvironmentRequestFactory(requestBody)
        return {
            :calledEndpoint => @@base_environment_endpoint,
            :sentValue => requestBody
        }
    end

    def self.listDatabaseServersResponseFactory(responseBody)
        return {
            :calledEndpoint => @@base_beams_dbserver_endpoint,
            :receivedValue => responseBody
        }
    end

    def self.getDatabaseServerByCrnResponseFactory(responseBody, crn)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/#{crn}",
            :receivedValue => responseBody
        }
    end

    def self.getDatabaseServerByNameResponseFactory(responseBody, name)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/name/#{name}",
            :receivedValue => responseBody
        }
    end

    def self.createDatabaseServerRequestFactory(requestBody)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/managed",
            :sentValue => requestBody
        }
    end

    def self.releaseManagedDatabaseServerResponseFactory(responseBody, crn)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/#{crn}/release",
            :receivedValue => responseBody
        }
    end

    def self.registerDatabaseServerRequestFactory(requestBody)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/register",
            :sentValue => requestBody
        }
    end

    def self.deleteDatabaseServerByCrnResponseFactory(responseBody, crn)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/#{crn}",
            :receivedValue => responseBody
        }
    end

    def self.listDatabasesResponseFactory(responseBody)
        return {
            :calledEndpoint => @@base_beams_db_endpoint,
            :receivedValue => responseBody
        }
    end

    def self.getDatabaseByCrnResponseFactory(responseBody, crn)
        return {
            :calledEndpoint => "#{@@base_beams_db_endpoint}/#{crn}",
            :receivedValue => responseBody
        }
    end

    def self.getDatabaseByNameResponseFactory(responseBody, name)
        return {
            :calledEndpoint => "#{@@base_beams_db_endpoint}/name/#{name}",
            :receivedValue => responseBody
        }
    end

    def self.createDatabaseOnServerRequestFactory(requestBody)
        return {
            :calledEndpoint => "#{@@base_beams_dbserver_endpoint}/createDatabase",
            :sentValue => requestBody
        }
    end

    def self.registerDatabaseRequestFactory(requestBody)
        return {
            :calledEndpoint => "#{@@base_beams_db_endpoint}/register",
            :sentValue => requestBody
        }
    end

    def self.deleteDatabaseByCrnResponseFactory(responseBody, crn)
        return {
            :calledEndpoint => "#{@@base_beams_db_endpoint}/#{crn}",
            :receivedValue => responseBody
        }
    end

    def self.deleteDatabaseByNameResponseFactory(responseBody, name)
        return {
            :calledEndpoint => "#{@@base_beams_db_endpoint}/name/#{name}",
            :receivedValue => responseBody
        }
    end
end
