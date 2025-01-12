// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	externalNonContainerDatabaseManagementRepresentation = map[string]interface{}{
		"external_database_connector_id":     Representation{RepType: Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"external_non_container_database_id": Representation{RepType: Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
		"license_model":                      Representation{RepType: Required, Create: `BRING_YOUR_OWN_LICENSE`},
		"enable_management":                  Representation{RepType: Required, Create: `true`, Update: `false`},
	}

	ExternalNonContainerDatabaseManagementResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", Required, Create, externalNonContainerDatabaseRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", Required, Create, externalDatabaseConnectorRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseExternalNonContainerDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_external_non_container_database_management.test_external_non_container_database_management"
	resourceNonCDB := "oci_database_external_non_container_database.test_external_non_container_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+ExternalNonContainerDatabaseManagementResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Required, Create, externalNonContainerDatabaseManagementRepresentation), "database", "externalNonContainerDatabaseManagement", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Required, Create, externalNonContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Required, Create, externalNonContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies,
		},
		// verify Update (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Optional, Create, externalNonContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify Update (Disable Database Management)
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Optional, Update, externalNonContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_non_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseManagementResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_database_external_non_container_database_management", "test_external_non_container_database_management", Optional, Update, externalNonContainerDatabaseManagementRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceNonCDB, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
