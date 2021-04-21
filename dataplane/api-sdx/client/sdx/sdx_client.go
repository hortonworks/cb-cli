// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sdx API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sdx API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Advertisedruntimes lists advertised datalake versions
*/
func (a *Client) Advertisedruntimes(params *AdvertisedruntimesParams) (*AdvertisedruntimesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdvertisedruntimesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "advertisedruntimes",
		Method:             "GET",
		PathPattern:        "/sdx/advertisedruntimes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AdvertisedruntimesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AdvertisedruntimesOK), nil

}

/*
BackupDatabase backups the database backing datalake
*/
func (a *Client) BackupDatabase(params *BackupDatabaseParams) (*BackupDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDatabaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backupDatabase",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/backupDatabase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupDatabaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackupDatabaseOK), nil

}

/*
BackupDatabaseStatus gets the status of datalake database backup operation
*/
func (a *Client) BackupDatabaseStatus(params *BackupDatabaseStatusParams) (*BackupDatabaseStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDatabaseStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backupDatabaseStatus",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/backupDatabaseStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupDatabaseStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackupDatabaseStatusOK), nil

}

/*
BackupDatalake backups the datalake
*/
func (a *Client) BackupDatalake(params *BackupDatalakeParams) (*BackupDatalakeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDatalakeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backupDatalake",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/backupDatalake",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupDatalakeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackupDatalakeOK), nil

}

/*
BackupDatalakeStatus backups status of the datalake
*/
func (a *Client) BackupDatalakeStatus(params *BackupDatalakeStatusParams) (*BackupDatalakeStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBackupDatalakeStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "backupDatalakeStatus",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/backupDatalakeStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BackupDatalakeStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*BackupDatalakeStatusOK), nil

}

/*
CreateCustomSdx creates custom s d x cluster
*/
func (a *Client) CreateCustomSdx(params *CreateCustomSdxParams) (*CreateCustomSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCustomSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCustomSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/custom_image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCustomSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCustomSdxOK), nil

}

/*
CreateSdx creates s d x cluster
*/
func (a *Client) CreateSdx(params *CreateSdxParams) (*CreateSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSdxOK), nil

}

/*
DeleteSdx deletes s d x cluster
*/
func (a *Client) DeleteSdx(params *DeleteSdxParams) (*DeleteSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdx",
		Method:             "DELETE",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSdxOK), nil

}

/*
DeleteSdxByCrn deletes s d x cluster by crn
*/
func (a *Client) DeleteSdxByCrn(params *DeleteSdxByCrnParams) (*DeleteSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdxByCrn",
		Method:             "DELETE",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSdxByCrnOK), nil

}

/*
GetDatabaseServerByClusterCrn gets database server for s d x cluster by cluster crn
*/
func (a *Client) GetDatabaseServerByClusterCrn(params *GetDatabaseServerByClusterCrnParams) (*GetDatabaseServerByClusterCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatabaseServerByClusterCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDatabaseServerByClusterCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}/dbserver",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDatabaseServerByClusterCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDatabaseServerByClusterCrnOK), nil

}

/*
GetRangerCloudIdentitySyncStatus gets status of a ranger cloud identity sync
*/
func (a *Client) GetRangerCloudIdentitySyncStatus(params *GetRangerCloudIdentitySyncStatusParams) (*GetRangerCloudIdentitySyncStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRangerCloudIdentitySyncStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRangerCloudIdentitySyncStatus",
		Method:             "GET",
		PathPattern:        "/sdx/envcrn/{envCrn}/ranger_cloud_identity_sync_status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRangerCloudIdentitySyncStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRangerCloudIdentitySyncStatusOK), nil

}

/*
GetSdx gets s d x cluster
*/
func (a *Client) GetSdx(params *GetSdxParams) (*GetSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdx",
		Method:             "GET",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxOK), nil

}

/*
GetSdxByCrn gets s d x cluster by crn
*/
func (a *Client) GetSdxByCrn(params *GetSdxByCrnParams) (*GetSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByCrnOK), nil

}

/*
GetSdxByEnvCrn gets s d x cluster by environment crn
*/
func (a *Client) GetSdxByEnvCrn(params *GetSdxByEnvCrnParams) (*GetSdxByEnvCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByEnvCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByEnvCrn",
		Method:             "GET",
		PathPattern:        "/sdx/envcrn/{envCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByEnvCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByEnvCrnOK), nil

}

/*
GetSdxDetail gets s d x cluster detail
*/
func (a *Client) GetSdxDetail(params *GetSdxDetailParams) (*GetSdxDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetail",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailOK), nil

}

/*
GetSdxDetailByCrn gets s d x cluster detail by crn
*/
func (a *Client) GetSdxDetailByCrn(params *GetSdxDetailByCrnParams) (*GetSdxDetailByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetailByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailByCrnOK), nil

}

/*
ListSdx lists s d x clusters
*/
func (a *Client) ListSdx(params *ListSdxParams) (*ListSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSdx",
		Method:             "GET",
		PathPattern:        "/sdx/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSdxOK), nil

}

/*
RenewCertificateOnSdxByCrn renews certificate on s d x cluster by crn
*/
func (a *Client) RenewCertificateOnSdxByCrn(params *RenewCertificateOnSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenewCertificateOnSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "renewCertificateOnSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/renew_certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RenewCertificateOnSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
RepairSdxNode repairs an sdx node in the specified hostgroup
*/
func (a *Client) RepairSdxNode(params *RepairSdxNodeParams) (*RepairSdxNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNode",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairSdxNodeOK), nil

}

/*
RepairSdxNodeByCrn repairs an sdx node in the specified hostgroup by crn
*/
func (a *Client) RepairSdxNodeByCrn(params *RepairSdxNodeByCrnParams) (*RepairSdxNodeByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNodeByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairSdxNodeByCrnOK), nil

}

/*
RestoreDatabase restores the database backing datalake
*/
func (a *Client) RestoreDatabase(params *RestoreDatabaseParams) (*RestoreDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreDatabaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreDatabase",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/restoreDatabase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreDatabaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RestoreDatabaseOK), nil

}

/*
RestoreDatabaseStatus gets the status of datalake database restore operation
*/
func (a *Client) RestoreDatabaseStatus(params *RestoreDatabaseStatusParams) (*RestoreDatabaseStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreDatabaseStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreDatabaseStatus",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/restoreDatabaseStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreDatabaseStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RestoreDatabaseStatusOK), nil

}

/*
RetrySdx retries sdx
*/
func (a *Client) RetrySdx(params *RetrySdxParams) (*RetrySdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrySdxOK), nil

}

/*
RetrySdxByCrn retries sdx by crn
*/
func (a *Client) RetrySdxByCrn(params *RetrySdxByCrnParams) (*RetrySdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrySdxByCrnOK), nil

}

/*
RotateAutoTLSCertificatesByCrn rotates the certificates of the cluster
*/
func (a *Client) RotateAutoTLSCertificatesByCrn(params *RotateAutoTLSCertificatesByCrnParams) (*RotateAutoTLSCertificatesByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRotateAutoTLSCertificatesByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rotateAutoTlsCertificatesByCrn",
		Method:             "PUT",
		PathPattern:        "/sdx/crn/{crn}/rotate_autotls_certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RotateAutoTLSCertificatesByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RotateAutoTLSCertificatesByCrnOK), nil

}

/*
RotateAutoTLSCertificatesByName rotates the certificates of the cluster
*/
func (a *Client) RotateAutoTLSCertificatesByName(params *RotateAutoTLSCertificatesByNameParams) (*RotateAutoTLSCertificatesByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRotateAutoTLSCertificatesByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "rotateAutoTlsCertificatesByName",
		Method:             "PUT",
		PathPattern:        "/sdx/{name}/rotate_autotls_certificates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RotateAutoTLSCertificatesByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RotateAutoTLSCertificatesByNameOK), nil

}

/*
SetRangerCloudIdentityMapping sets ranger cloud identity mapping
*/
func (a *Client) SetRangerCloudIdentityMapping(params *SetRangerCloudIdentityMappingParams) (*SetRangerCloudIdentityMappingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRangerCloudIdentityMappingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setRangerCloudIdentityMapping",
		Method:             "POST",
		PathPattern:        "/sdx/envcrn/{envCrn}/ranger_cloud_identity_mapping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetRangerCloudIdentityMappingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetRangerCloudIdentityMappingOK), nil

}

/*
StartSdxByCrn starts sdx by crn
*/
func (a *Client) StartSdxByCrn(params *StartSdxByCrnParams) (*StartSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartSdxByCrnOK), nil

}

/*
StartSdxByName starts sdx
*/
func (a *Client) StartSdxByName(params *StartSdxByNameParams) (*StartSdxByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartSdxByNameOK), nil

}

/*
StopSdxByCrn stops sdx by crn
*/
func (a *Client) StopSdxByCrn(params *StopSdxByCrnParams) (*StopSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopSdxByCrnOK), nil

}

/*
StopSdxByName stops sdx
*/
func (a *Client) StopSdxByName(params *StopSdxByNameParams) (*StopSdxByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopSdxByNameOK), nil

}

/*
SyncSdx syncs s d x cluster by name
*/
func (a *Client) SyncSdx(params *SyncSdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncSdxByCrn syncs s d x cluster by crn
*/
func (a *Client) SyncSdxByCrn(params *SyncSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpgradeDatalakeCluster upgrades the datalake cluster
*/
func (a *Client) UpgradeDatalakeCluster(params *UpgradeDatalakeClusterParams) (*UpgradeDatalakeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeDatalakeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeDatalakeCluster",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeDatalakeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeDatalakeClusterOK), nil

}

/*
UpgradeDatalakeClusterByCrn upgrades the datalake cluster
*/
func (a *Client) UpgradeDatalakeClusterByCrn(params *UpgradeDatalakeClusterByCrnParams) (*UpgradeDatalakeClusterByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeDatalakeClusterByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeDatalakeClusterByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeDatalakeClusterByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeDatalakeClusterByCrnOK), nil

}

/*
Versions lists datalake versions
*/
func (a *Client) Versions(params *VersionsParams) (*VersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "versions",
		Method:             "GET",
		PathPattern:        "/sdx/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
