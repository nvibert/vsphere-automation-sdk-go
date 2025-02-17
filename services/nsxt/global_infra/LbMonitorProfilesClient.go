/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: LbMonitorProfiles
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
)

type LbMonitorProfilesClient interface {

    // Delete the LBMonitorProfile along with all the entities contained by this LBMonitorProfile.
    //
    // @param lbMonitorProfileIdParam LBMonitorProfile ID (required)
    // @param forceParam Force delete the resource even if it is being used somewhere (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(lbMonitorProfileIdParam string, forceParam *bool) error

    // Read a LBMonitorProfile.
    //
    // @param lbMonitorProfileIdParam LBMonitorProfile ID (required)
    // @return com.vmware.nsx_policy.model.LBMonitorProfile
    // The return value will contain all the properties defined in model.LBMonitorProfile.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(lbMonitorProfileIdParam string) (*data.StructValue, error)

    // Paginated list of all LBMonitorProfiles for infra.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.LBMonitorProfileListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.LBMonitorProfileListResult, error)

    // If a LBMonitorProfile with the lb-monitor-profile-id is not already present, create a new LBMonitorProfile. If it already exists, update the LBMonitorProfile. This is a full replace.
    //
    // @param lbMonitorProfileIdParam LBMonitorProfile ID (required)
    // @param lbMonitorProfileParam (required)
    // The parameter must contain all the properties defined in model.LBMonitorProfile.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(lbMonitorProfileIdParam string, lbMonitorProfileParam *data.StructValue) error

    // If a LBMonitorProfile with the lb-monitor-profile-id is not already present, create a new LBMonitorProfile. If it already exists, update the LBMonitorProfile. This is a full replace.
    //
    // @param lbMonitorProfileIdParam LBMonitorProfile ID (required)
    // @param lbMonitorProfileParam (required)
    // The parameter must contain all the properties defined in model.LBMonitorProfile.
    // @return com.vmware.nsx_policy.model.LBMonitorProfile
    // The return value will contain all the properties defined in model.LBMonitorProfile.
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(lbMonitorProfileIdParam string, lbMonitorProfileParam *data.StructValue) (*data.StructValue, error)
}
