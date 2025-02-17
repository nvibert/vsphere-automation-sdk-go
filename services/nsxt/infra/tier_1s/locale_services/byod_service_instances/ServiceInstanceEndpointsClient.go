/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: ServiceInstanceEndpoints
 * Used by client-side stubs.
 */

package byod_service_instances

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type ServiceInstanceEndpointsClient interface {

    // Delete Tier1 service instance endpoint
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @param serviceInstanceEndpointIdParam Tier1 Service instance endpoint id (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, serviceInstanceEndpointIdParam string) error

    // Read Tier1 service instance endpoint
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @param serviceInstanceEndpointIdParam Tier1 Service instance endpoint id (required)
    // @return com.vmware.nsx_policy.model.ServiceInstanceEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, serviceInstanceEndpointIdParam string) (model.ServiceInstanceEndpoint, error)

    // List all Tier1 service instance endpoint
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ServiceInstanceEndpointListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ServiceInstanceEndpointListResult, error)

    // Create Tier1 Service instance endpoint.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @param serviceInstanceEndpointIdParam Tier1 Service instance endpoint id (required)
    // @param serviceInstanceEndpointParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, serviceInstanceEndpointIdParam string, serviceInstanceEndpointParam model.ServiceInstanceEndpoint) error

    // Create Tier1 service instance endpoint with given request if not exist. Modification of Tier1 service instance endpoint is not allowed.
    //
    // @param tier1IdParam Tier-1 id (required)
    // @param localeServiceIdParam Locale service id (required)
    // @param serviceInstanceIdParam Service instance id (required)
    // @param serviceInstanceEndpointIdParam Tier1 Service instance endpoint id (required)
    // @param serviceInstanceEndpointParam (required)
    // @return com.vmware.nsx_policy.model.ServiceInstanceEndpoint
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(tier1IdParam string, localeServiceIdParam string, serviceInstanceIdParam string, serviceInstanceEndpointIdParam string, serviceInstanceEndpointParam model.ServiceInstanceEndpoint) (model.ServiceInstanceEndpoint, error)
}
