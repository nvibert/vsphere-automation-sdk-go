/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Interface file for service: SpoofguardProfiles
 * Used by client-side stubs.
 */

package global_infra

import (
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
)

type SpoofguardProfilesClient interface {

    // API will delete SpoofGuard profile with the given id.
    //
    // @param spoofguardProfileIdParam SpoofGuard profile id (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Delete(spoofguardProfileIdParam string, overrideParam *bool) error

    // API will return details of the SpoofGuard profile with given id. If the profile does not exist, it will return 404.
    //
    // @param spoofguardProfileIdParam SpoofGuard profile id (required)
    // @return com.vmware.nsx_policy.model.SpoofGuardProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Get(spoofguardProfileIdParam string) (model.SpoofGuardProfile, error)

    // API will list all SpoofGuard profiles.
    //
    // @param cursorParam Opaque cursor to be used for getting next page of records (supplied by current result page) (optional)
    // @param includeMarkForDeleteObjectsParam Include objects that are marked for deletion in results (optional, default to false)
    // @param includedFieldsParam Comma separated list of fields that should be included in query result (optional)
    // @param pageSizeParam Maximum number of results to return in this page (server may return fewer) (optional, default to 1000)
    // @param sortAscendingParam (optional)
    // @param sortByParam Field by which records are sorted (optional)
    // @return com.vmware.nsx_policy.model.SpoofGuardProfileListResult
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	List(cursorParam *string, includeMarkForDeleteObjectsParam *bool, includedFieldsParam *string, pageSizeParam *int64, sortAscendingParam *bool, sortByParam *string) (model.SpoofGuardProfileListResult, error)

    // Create a new SpoofGuard profile if the SpoofGuard profile with the given id does not exist. Otherwise, patch with the existing SpoofGuard profile.
    //
    // @param spoofguardProfileIdParam SpoofGuard profile id (required)
    // @param spoofGuardProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Patch(spoofguardProfileIdParam string, spoofGuardProfileParam model.SpoofGuardProfile, overrideParam *bool) error

    // API will create or replace SpoofGuard profile.
    //
    // @param spoofguardProfileIdParam SpoofGuard profile id (required)
    // @param spoofGuardProfileParam (required)
    // @param overrideParam Locally override the global object (optional, default to false)
    // @return com.vmware.nsx_policy.model.SpoofGuardProfile
    // @throws InvalidRequest  Bad Request, Precondition Failed
    // @throws Unauthorized  Forbidden
    // @throws ServiceUnavailable  Service Unavailable
    // @throws InternalServerError  Internal Server Error
    // @throws NotFound  Not Found
	Update(spoofguardProfileIdParam string, spoofGuardProfileParam model.SpoofGuardProfile, overrideParam *bool) (model.SpoofGuardProfile, error)
}
