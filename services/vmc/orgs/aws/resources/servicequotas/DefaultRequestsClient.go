
/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Client stubs for service: Requests
 * Functions that implement the generated RequestsClient interface
 */


package servicequotas

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/model"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

type DefaultRequestsClient struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodIdentifiers   []core.MethodIdentifier
	methodNameToDefMap  map[string]*core.MethodDefinition
	errorBindingMap     map[string]bindings.BindingType
	interfaceIdentifier core.InterfaceIdentifier
	connector           client.Connector
}

func NewDefaultRequestsClient(connector client.Connector) *DefaultRequestsClient {
	interfaceName := "com.vmware.vmc.orgs.aws.resources.servicequotas.requests"
	interfaceIdentifier := core.NewInterfaceIdentifier(interfaceName)
	methodIdentifiers := []core.MethodIdentifier{
		core.NewMethodIdentifier(interfaceIdentifier, "get"),
		core.NewMethodIdentifier(interfaceIdentifier, "list"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorBindingMap := make(map[string]bindings.BindingType)
	errorBindingMap[errors.AlreadyExists{}.Error()] = errors.AlreadyExistsBindingType()
	errorBindingMap[errors.AlreadyInDesiredState{}.Error()] = errors.AlreadyInDesiredStateBindingType()
	errorBindingMap[errors.Canceled{}.Error()] = errors.CanceledBindingType()
	errorBindingMap[errors.ConcurrentChange{}.Error()] = errors.ConcurrentChangeBindingType()
	errorBindingMap[errors.Error{}.Error()] = errors.ErrorBindingType()
	errorBindingMap[errors.FeatureInUse{}.Error()] = errors.FeatureInUseBindingType()
	errorBindingMap[errors.InternalServerError{}.Error()] = errors.InternalServerErrorBindingType()
	errorBindingMap[errors.InvalidArgument{}.Error()] = errors.InvalidArgumentBindingType()
	errorBindingMap[errors.InvalidElementConfiguration{}.Error()] = errors.InvalidElementConfigurationBindingType()
	errorBindingMap[errors.InvalidElementType{}.Error()] = errors.InvalidElementTypeBindingType()
	errorBindingMap[errors.InvalidRequest{}.Error()] = errors.InvalidRequestBindingType()
	errorBindingMap[errors.NotFound{}.Error()] = errors.NotFoundBindingType()
	errorBindingMap[errors.NotAllowedInCurrentState{}.Error()] = errors.NotAllowedInCurrentStateBindingType()
	errorBindingMap[errors.OperationNotFound{}.Error()] = errors.OperationNotFoundBindingType()
	errorBindingMap[errors.ResourceBusy{}.Error()] = errors.ResourceBusyBindingType()
	errorBindingMap[errors.ResourceInUse{}.Error()] = errors.ResourceInUseBindingType()
	errorBindingMap[errors.ResourceInaccessible{}.Error()] = errors.ResourceInaccessibleBindingType()
	errorBindingMap[errors.ServiceUnavailable{}.Error()] = errors.ServiceUnavailableBindingType()
	errorBindingMap[errors.TimedOut{}.Error()] = errors.TimedOutBindingType()
	errorBindingMap[errors.UnableToAllocateResource{}.Error()] = errors.UnableToAllocateResourceBindingType()
	errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errorBindingMap[errors.UnexpectedInput{}.Error()] = errors.UnexpectedInputBindingType()
	errorBindingMap[errors.Unsupported{}.Error()] = errors.UnsupportedBindingType()
	errorBindingMap[errors.UnverifiedPeer{}.Error()] = errors.UnverifiedPeerBindingType()


	rIface := DefaultRequestsClient{interfaceName: interfaceName, methodIdentifiers: methodIdentifiers, interfaceDefinition: interfaceDefinition, errorBindingMap: errorBindingMap, interfaceIdentifier: interfaceIdentifier, connector: connector}
	rIface.methodNameToDefMap = make(map[string]*core.MethodDefinition)
	rIface.methodNameToDefMap["get"] = rIface.getMethodDefinition()
	rIface.methodNameToDefMap["list"] = rIface.listMethodDefinition()
	return &rIface
}

func (rIface *DefaultRequestsClient) Get(orgParam string, serviceQuotaRequestIdParam string) (model.ServiceQuotaRequest, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "get")
	sv := bindings.NewStructValueBuilder(requestsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("ServiceQuotaRequestId", serviceQuotaRequestIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput model.ServiceQuotaRequest
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := requestsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput model.ServiceQuotaRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), requestsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.(model.ServiceQuotaRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}

func (rIface *DefaultRequestsClient) List(orgParam string) ([]model.ServiceQuotaRequest, error) {
	typeConverter := rIface.connector.TypeConverter()
	methodIdentifier := core.NewMethodIdentifier(rIface.interfaceIdentifier, "list")
	sv := bindings.NewStructValueBuilder(requestsListInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.ServiceQuotaRequest
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := requestsListRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	executionContext := rIface.connector.NewExecutionContext()
	methodResult := rIface.Invoke(executionContext, methodIdentifier, inputDataValue)
	var emptyOutput []model.ServiceQuotaRequest
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), requestsListOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.ServiceQuotaRequest), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.errorBindingMap[methodResult.Error().Name()])
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}


func (rIface *DefaultRequestsClient) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, inputDataValue data.DataValue) core.MethodResult {
	methodResult := rIface.connector.GetApiProvider().Invoke(rIface.interfaceName, methodId.Name(), inputDataValue, ctx)
	return methodResult
}


func (rIface *DefaultRequestsClient) getMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(requestsGetInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(requestsGetOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.get method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.get method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "get")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.get method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.get method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}

func (rIface *DefaultRequestsClient) listMethodDefinition() *core.MethodDefinition {
	interfaceIdentifier := core.NewInterfaceIdentifier(rIface.interfaceName)
	typeConverter := rIface.connector.TypeConverter()

	input, inputError := typeConverter.ConvertToDataDefinition(requestsListInputType())
	output, outputError := typeConverter.ConvertToDataDefinition(requestsListOutputType())
	if inputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.list method's input - %s",
			bindings.VAPIerrorsToError(inputError).Error())
		return nil
	}
	if outputError != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.list method's output - %s",
			bindings.VAPIerrorsToError(outputError).Error())
		return nil
	}
	methodIdentifier := core.NewMethodIdentifier(interfaceIdentifier, "list")
	errorDefinitions := make([]data.ErrorDefinition, 0)
	rIface.errorBindingMap[errors.Unauthenticated{}.Error()] = errors.UnauthenticatedBindingType()
	errDef1, errError1 := typeConverter.ConvertToDataDefinition(errors.UnauthenticatedBindingType())
	if errError1 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.list method's errors.Unauthenticated error - %s",
			bindings.VAPIerrorsToError(errError1).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef1.(data.ErrorDefinition))
	rIface.errorBindingMap[errors.Unauthorized{}.Error()] = errors.UnauthorizedBindingType()
	errDef2, errError2 := typeConverter.ConvertToDataDefinition(errors.UnauthorizedBindingType())
	if errError2 != nil {
		log.Errorf("Error in ConvertToDataDefinition for DefaultRequestsClient.list method's errors.Unauthorized error - %s",
			bindings.VAPIerrorsToError(errError2).Error())
		return nil
	}
	errorDefinitions = append(errorDefinitions, errDef2.(data.ErrorDefinition))

	methodDefinition := core.NewMethodDefinition(methodIdentifier, input, output, errorDefinitions)
	return &methodDefinition
}
