/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: GroupAssociations
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type GroupAssociationsClient interface {

    // Get policy groups for which the given object is a member. In Federation environment, if the given object is a global entity (eg: global segment) and if the entity is not stretched to the site specified in the enforcement_point_path parameter,then the following is returned:- 1)If the entity is a member of any global group and that group is stretched to the enforcement_point_path site,then the API returns an empty list. 2)If the entity is not a member of any global group,this API returns an 'invalid path' error message. 3)If both the entity and its corresponding groups are stretched to the enforcement_point_path site , then the API returns the groups list.
    //
    // @param intentPathParam String path of the intent object (required)
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param enforcementPointPathParam String Path of the enforcement point (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.PolicyResourceReferenceForEPListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(intentPathParam string, cursorParam *string, enforcementPointPathParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.PolicyResourceReferenceForEPListResult, error)
}
