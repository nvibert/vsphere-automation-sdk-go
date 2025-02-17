/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: FirewallFloodProtectionProfileBindingMaps.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package groups

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt-gm/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)





func firewallFloodProtectionProfileBindingMapsDeleteInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallFloodProtectionProfileBindingMapsDeleteOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallFloodProtectionProfileBindingMapsDeleteRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = bindings.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"DELETE",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallFloodProtectionProfileBindingMapsGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
}

func firewallFloodProtectionProfileBindingMapsGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = bindings.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsListInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallFloodProtectionProfileBindingMapsListOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapListResultBindingType)
}

func firewallFloodProtectionProfileBindingMapsListRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	fields["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fields["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fields["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["cursor"] = "Cursor"
	fieldNameMap["include_mark_for_delete_objects"] = "IncludeMarkForDeleteObjects"
	fieldNameMap["included_fields"] = "IncludedFields"
	fieldNameMap["page_size"] = "PageSize"
	fieldNameMap["sort_ascending"] = "SortAscending"
	fieldNameMap["sort_by"] = "SortBy"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["included_fields"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["page_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	paramsTypeMap["include_mark_for_delete_objects"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["cursor"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_by"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["sort_ascending"] = bindings.NewOptionalType(bindings.NewBooleanType())
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
	queryParams["cursor"] = "cursor"
	queryParams["sort_ascending"] = "sort_ascending"
	queryParams["included_fields"] = "included_fields"
	queryParams["sort_by"] = "sort_by"
	queryParams["include_mark_for_delete_objects"] = "include_mark_for_delete_objects"
	queryParams["page_size"] = "page_size"
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
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsPatchInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallFloodProtectionProfileBindingMapsPatchOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func firewallFloodProtectionProfileBindingMapsPatchRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = bindings.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_flood_protection_profile_binding_map",
		"PATCH",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		204,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}

func firewallFloodProtectionProfileBindingMapsUpdateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallFloodProtectionProfileBindingMapsUpdateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
}

func firewallFloodProtectionProfileBindingMapsUpdateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["domain_id"] = bindings.NewStringType()
	fields["group_id"] = bindings.NewStringType()
	fields["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	fields["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	fieldNameMap["domain_id"] = "DomainId"
	fieldNameMap["group_id"] = "GroupId"
	fieldNameMap["firewall_flood_protection_profile_binding_map_id"] = "FirewallFloodProtectionProfileBindingMapId"
	fieldNameMap["policy_firewall_flood_protection_profile_binding_map"] = "PolicyFirewallFloodProtectionProfileBindingMap"
	paramsTypeMap["domain_id"] = bindings.NewStringType()
	paramsTypeMap["group_id"] = bindings.NewStringType()
	paramsTypeMap["policy_firewall_flood_protection_profile_binding_map"] = bindings.NewReferenceType(model.PolicyFirewallFloodProtectionProfileBindingMapBindingType)
	paramsTypeMap["firewall_flood_protection_profile_binding_map_id"] = bindings.NewStringType()
	paramsTypeMap["domainId"] = bindings.NewStringType()
	paramsTypeMap["groupId"] = bindings.NewStringType()
	paramsTypeMap["firewallFloodProtectionProfileBindingMapId"] = bindings.NewStringType()
	pathParams["firewall_flood_protection_profile_binding_map_id"] = "firewallFloodProtectionProfileBindingMapId"
	pathParams["group_id"] = "groupId"
	pathParams["domain_id"] = "domainId"
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
		"policy_firewall_flood_protection_profile_binding_map",
		"PUT",
		"/global-manager/api/v1/global-infra/domains/{domainId}/groups/{groupId}/firewall-flood-protection-profile-binding-maps/{firewallFloodProtectionProfileBindingMapId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


