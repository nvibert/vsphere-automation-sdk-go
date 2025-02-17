/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: FirewallIdentityStoreLdapServer.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package global_infra

import (
	"reflect"
	"github.com/vmware/vsphere-automation-sdk-go/services/nsxt/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
)

// Possible value for ``action`` of method FirewallIdentityStoreLdapServer#create.
const FirewallIdentityStoreLdapServer_CREATE_ACTION_CONNECTIVITY = "CONNECTIVITY"




func firewallIdentityStoreLdapServerCreateInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fields["action"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	fieldNameMap["action"] = "Action"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func firewallIdentityStoreLdapServerCreateOutputType() bindings.BindingType {
	return bindings.NewReferenceType(model.DirectoryLdapServerStatusBindingType)
}

func firewallIdentityStoreLdapServerCreateRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	fields["action"] = bindings.NewStringType()
	fields["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["directory_ldap_server"] = "DirectoryLdapServer"
	fieldNameMap["action"] = "Action"
	fieldNameMap["enforcement_point_path"] = "EnforcementPointPath"
	paramsTypeMap["enforcement_point_path"] = bindings.NewOptionalType(bindings.NewStringType())
	paramsTypeMap["action"] = bindings.NewStringType()
	paramsTypeMap["directory_ldap_server"] = bindings.NewReferenceType(model.DirectoryLdapServerBindingType)
	queryParams["action"] = "action"
	queryParams["enforcement_point_path"] = "enforcement_point_path"
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
		"directory_ldap_server",
		"POST",
		"/policy/api/v1/global-infra/firewall-identity-store-ldap-server",
		"",
		resultHeaders,
		201,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400,"com.vmware.vapi.std.errors.unauthorized": 403,"com.vmware.vapi.std.errors.service_unavailable": 503,"com.vmware.vapi.std.errors.internal_server_error": 500,"com.vmware.vapi.std.errors.not_found": 404})
}


