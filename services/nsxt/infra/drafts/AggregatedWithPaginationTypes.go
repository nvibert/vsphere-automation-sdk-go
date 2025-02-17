/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Aggregated_with_pagination.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package drafts

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func aggregatedWithPaginationGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["draft_id"] = bindings.NewStringType()
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["request_id"] = "RequestId"
	fieldNameMap["root_path"] = "RootPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func aggregatedWithPaginationGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyDraftPaginatedAggregatedConfigurationResultBindingType)
}

func aggregatedWithPaginationGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["draft_id"] = bindings.NewStringType()
	fields["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["draft_id"] = "DraftId"
	fieldNameMap["request_id"] = "RequestId"
	fieldNameMap["root_path"] = "RootPath"
	paramsTypeMap["request_id"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["root_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["draft_id"] = bindings.NewStringType()
	paramsTypeMap["draftId"] = bindings.NewStringType()
	pathParams["draft_id"] = "draftId"
	queryParams["root_path"] = "root_path"
	queryParams["request_id"] = "request_id"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/policy/api/v1/infra/drafts/{draftId}/aggregated_with_pagination",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


