/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: State.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package dhcp_static_bindings

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func stateGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func stateGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DhcpStaticBindingStateBindingType)
}

func stateGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["tier1_id"] = bindings.NewStringType()
	fields["segment_id"] = bindings.NewStringType()
	fields["binding_id"] = bindings.NewStringType()
	fieldNameMap["tier1_id"] = "Tier1Id"
	fieldNameMap["segment_id"] = "SegmentId"
	fieldNameMap["binding_id"] = "BindingId"
	paramsTypeMap["tier1_id"] = bindings.NewStringType()
	paramsTypeMap["segment_id"] = bindings.NewStringType()
	paramsTypeMap["binding_id"] = bindings.NewStringType()
	paramsTypeMap["tier1Id"] = bindings.NewStringType()
	paramsTypeMap["segmentId"] = bindings.NewStringType()
	paramsTypeMap["bindingId"] = bindings.NewStringType()
	pathParams["segment_id"] = "segmentId"
	pathParams["tier1_id"] = "tier1Id"
	pathParams["binding_id"] = "bindingId"
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
		"/policy/api/v1/global-infra/tier-1s/{tier1Id}/segments/{segmentId}/dhcp-static-bindings/{bindingId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


