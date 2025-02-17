/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: IpDiscoveryProfiles
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type IpDiscoveryProfilesClient interface {

    // API will delete IP Discovery profile.
    //
    // @param ipDiscoveryProfileIdParam IP Discovery Profile ID (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(ipDiscoveryProfileIdParam string, overrideParam *bool) error

    // API will get IP Discovery profile.
    //
    // @param ipDiscoveryProfileIdParam IP Discovery Profile ID (required)
    // @return com.vmware.nsx_policy.model.IPDiscoveryProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(ipDiscoveryProfileIdParam string) (model.IPDiscoveryProfile, error)

    // API will list all IP Discovery Profiles active in current discovery profile id.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.IPDiscoveryProfileListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.IPDiscoveryProfileListResult, error)

    // API will create IP Discovery profile.
    //
    // @param ipDiscoveryProfileIdParam IP Discovery Profile ID (required)
    // @param ipDiscoveryProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(ipDiscoveryProfileIdParam string, ipDiscoveryProfileParam model.IPDiscoveryProfile, overrideParam *bool) error

    // API will update IP Discovery profile.
    //
    // @param ipDiscoveryProfileIdParam IP Discovery Profile ID (required)
    // @param ipDiscoveryProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @return com.vmware.nsx_policy.model.IPDiscoveryProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(ipDiscoveryProfileIdParam string, ipDiscoveryProfileParam model.IPDiscoveryProfile, overrideParam *bool) (model.IPDiscoveryProfile, error)
}
