/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: AlbL4PolicySets
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type AlbL4PolicySetsClient interface {

    // Delete the ALBL4PolicySet along with all the entities contained by this ALBL4PolicySet.
    //
    // @param albL4policysetIdParam ALBL4PolicySet ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(albL4policysetIdParam string, forceParam *bool) error

    // Read a ALBL4PolicySet.
    //
    // @param albL4policysetIdParam ALBL4PolicySet ID (required)
    // @return com.vmware.nsx_policy.model.ALBL4PolicySet
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(albL4policysetIdParam string) (model.ALBL4PolicySet, error)

    // Paginated list of all ALBL4PolicySet for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.ALBL4PolicySetApiResponse
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.ALBL4PolicySetApiResponse, error)

    // If a ALBl4policyset with the alb-l4policyset-id is not already present, create a new ALBl4policyset. If it already exists, update the ALBl4policyset. This is a full replace.
    //
    // @param albL4policysetIdParam ALBl4policyset ID (required)
    // @param aLBL4PolicySetParam (required)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(albL4policysetIdParam string, aLBL4PolicySetParam model.ALBL4PolicySet) error

    // If a ALBL4PolicySet with the alb-L4PolicySet-id is not already present, create a new ALBL4PolicySet. If it already exists, update the ALBL4PolicySet. This is a full replace.
    //
    // @param albL4policysetIdParam ALBL4PolicySet ID (required)
    // @param aLBL4PolicySetParam (required)
    // @return com.vmware.nsx_policy.model.ALBL4PolicySet
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(albL4policysetIdParam string, aLBL4PolicySetParam model.ALBL4PolicySet) (model.ALBL4PolicySet, error)
}
