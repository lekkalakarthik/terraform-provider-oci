// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_database "github.com/oracle/oci-go-sdk/v39/database"
	oci_work_requests "github.com/oracle/oci-go-sdk/v39/workrequests"
)

func init() {
	RegisterResource("oci_database_external_pluggable_database_management", DatabaseExternalPluggableDatabaseManagementResource())
}

func DatabaseExternalPluggableDatabaseManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createDatabaseExternalPluggableDatabaseManagement,
		Update:   updateDatabaseExternalPluggableDatabaseManagement,
		Read:     readDatabaseExternalPluggableDatabaseManagement,
		Delete:   deleteDatabaseExternalPluggableDatabaseManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"external_database_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"external_pluggable_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_management": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Computed
		},
	}
}

func createDatabaseExternalPluggableDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalPluggableDatabaseManagementResponse{}
	return CreateResource(d, sync)
}

func updateDatabaseExternalPluggableDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExternalPluggableDatabaseManagementResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()
	sync.workRequestClient = m.(*OracleClients).workRequestClient
	sync.Res = &DatabaseExternalPluggableDatabaseManagementResponse{}
	return UpdateResource(d, sync)
}

func readDatabaseExternalPluggableDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteDatabaseExternalPluggableDatabaseManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DatabaseExternalPluggableDatabaseManagementResponse struct {
	enableResponse  *oci_database.EnableExternalPluggableDatabaseDatabaseManagementResponse
	disableResponse *oci_database.DisableExternalPluggableDatabaseDatabaseManagementResponse
}

type DatabaseExternalPluggableDatabaseManagementResourceCrud struct {
	BaseCrud
	Client                 *oci_database.DatabaseClient
	workRequestClient      *oci_work_requests.WorkRequestClient
	Res                    *DatabaseExternalPluggableDatabaseManagementResponse
	DisableNotFoundRetries bool
}

func (s *DatabaseExternalPluggableDatabaseManagementResourceCrud) ID() string {
	return GenerateDataSourceHashID("DatabaseExternalPluggableDatabaseManagementResource-", DatabaseExternalPluggableDatabaseManagementResource(), s.D)
}

func (s *DatabaseExternalPluggableDatabaseManagementResourceCrud) Create() error {

	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}

	if operation {
		// Enable Database Management
		request := oci_database.EnableExternalPluggableDatabaseDatabaseManagementRequest{}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalPluggableDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalPluggableDatabaseDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}
	// Disable Database Management
	request := oci_database.DisableExternalPluggableDatabaseDatabaseManagementRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalPluggableDatabaseManagementResourceCrud) Update() error {
	var operation bool
	if enableManagement, ok := s.D.GetOkExists("enable_management"); ok {
		operation = enableManagement.(bool)
	}
	operation = false
	if operation {
		// Enable Database Management
		request := oci_database.EnableExternalPluggableDatabaseDatabaseManagementRequest{}

		if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
			tmp := externalPluggableDatabaseId.(string)
			request.ExternalPluggableDatabaseId = &tmp
		}

		if externalDatabaseConnectorId, ok := s.D.GetOkExists("external_database_connector_id"); ok {
			tmp := externalDatabaseConnectorId.(string)
			request.EnableExternalPluggableDatabaseDatabaseManagementDetails.ExternalDatabaseConnectorId = &tmp
		}

		request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

		response, err := s.Client.EnableExternalPluggableDatabaseDatabaseManagement(context.Background(), request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		if workId != nil {
			_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
			if err != nil {
				return err
			}
		}
		s.Res.enableResponse = &response
		return nil
	}
	// Disable Database Management
	request := oci_database.DisableExternalPluggableDatabaseDatabaseManagementRequest{}

	if externalPluggableDatabaseId, ok := s.D.GetOkExists("external_pluggable_database_id"); ok {
		tmp := externalPluggableDatabaseId.(string)
		request.ExternalPluggableDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.DisableExternalPluggableDatabaseDatabaseManagement(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		_, err = WaitForWorkRequestWithErrorHandling(s.workRequestClient, workId, "externalPluggableDatabase", oci_work_requests.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate), s.DisableNotFoundRetries)
		if err != nil {
			return err
		}
	}
	s.Res.disableResponse = &response
	return nil
}

func (s *DatabaseExternalPluggableDatabaseManagementResourceCrud) SetData() error {
	return nil
}
