# Go API client for iwfidl

This APIs for iwf SDKs to operate workflows

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import iwfidl "github.com/indeedeng/iwf-idl"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), iwfidl.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), iwfidl.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), iwfidl.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), iwfidl.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://petstore.swagger.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultAPI* | [**ApiV1WorkflowConfigUpdatePost**](docs/DefaultAPI.md#apiv1workflowconfigupdatepost) | **Post** /api/v1/workflow/config/update | update the config of a workflow
*DefaultAPI* | [**ApiV1WorkflowDataobjectsGetPost**](docs/DefaultAPI.md#apiv1workflowdataobjectsgetpost) | **Post** /api/v1/workflow/dataobjects/get | get workflow data objects
*DefaultAPI* | [**ApiV1WorkflowGetPost**](docs/DefaultAPI.md#apiv1workflowgetpost) | **Post** /api/v1/workflow/get | get a workflow&#39;s status and results(if completed &amp; requested)
*DefaultAPI* | [**ApiV1WorkflowGetWithWaitPost**](docs/DefaultAPI.md#apiv1workflowgetwithwaitpost) | **Post** /api/v1/workflow/getWithWait | get a workflow&#39;s status and results(if completed &amp; requested), wait if the workflow is still running
*DefaultAPI* | [**ApiV1WorkflowInternalDumpPost**](docs/DefaultAPI.md#apiv1workflowinternaldumppost) | **Post** /api/v1/workflow/internal/dump | dump internal info of a workflow
*DefaultAPI* | [**ApiV1WorkflowResetPost**](docs/DefaultAPI.md#apiv1workflowresetpost) | **Post** /api/v1/workflow/reset | reset a workflow
*DefaultAPI* | [**ApiV1WorkflowRpcPost**](docs/DefaultAPI.md#apiv1workflowrpcpost) | **Post** /api/v1/workflow/rpc | execute an RPC of a workflow
*DefaultAPI* | [**ApiV1WorkflowSearchPost**](docs/DefaultAPI.md#apiv1workflowsearchpost) | **Post** /api/v1/workflow/search | search for workflows by a search attribute query
*DefaultAPI* | [**ApiV1WorkflowSearchattributesGetPost**](docs/DefaultAPI.md#apiv1workflowsearchattributesgetpost) | **Post** /api/v1/workflow/searchattributes/get | get workflow search attributes
*DefaultAPI* | [**ApiV1WorkflowSignalPost**](docs/DefaultAPI.md#apiv1workflowsignalpost) | **Post** /api/v1/workflow/signal | signal a workflow
*DefaultAPI* | [**ApiV1WorkflowStartPost**](docs/DefaultAPI.md#apiv1workflowstartpost) | **Post** /api/v1/workflow/start | start a workflow
*DefaultAPI* | [**ApiV1WorkflowStateDecidePost**](docs/DefaultAPI.md#apiv1workflowstatedecidepost) | **Post** /api/v1/workflowState/decide | for invoking WorkflowState.execute API
*DefaultAPI* | [**ApiV1WorkflowStateStartPost**](docs/DefaultAPI.md#apiv1workflowstatestartpost) | **Post** /api/v1/workflowState/start | for invoking WorkflowState.waitUntil API
*DefaultAPI* | [**ApiV1WorkflowStopPost**](docs/DefaultAPI.md#apiv1workflowstoppost) | **Post** /api/v1/workflow/stop | stop a workflow
*DefaultAPI* | [**ApiV1WorkflowTimerSkipPost**](docs/DefaultAPI.md#apiv1workflowtimerskippost) | **Post** /api/v1/workflow/timer/skip | skip the timer of a workflow
*DefaultAPI* | [**ApiV1WorkflowWaitForStateCompletionPost**](docs/DefaultAPI.md#apiv1workflowwaitforstatecompletionpost) | **Post** /api/v1/workflow/waitForStateCompletion | 
*DefaultAPI* | [**ApiV1WorkflowWorkerRpcPost**](docs/DefaultAPI.md#apiv1workflowworkerrpcpost) | **Post** /api/v1/workflowWorker/rpc | for invoking workflow RPC API in the worker
*DefaultAPI* | [**InfoHealthcheckGet**](docs/DefaultAPI.md#infohealthcheckget) | **Get** /info/healthcheck | return health info of the server


## Documentation For Models

 - [ChannelRequestStatus](docs/ChannelRequestStatus.md)
 - [CommandCombination](docs/CommandCombination.md)
 - [CommandRequest](docs/CommandRequest.md)
 - [CommandResults](docs/CommandResults.md)
 - [CommandWaitingType](docs/CommandWaitingType.md)
 - [Context](docs/Context.md)
 - [EncodedObject](docs/EncodedObject.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [ErrorSubStatus](docs/ErrorSubStatus.md)
 - [ExecuteApiFailurePolicy](docs/ExecuteApiFailurePolicy.md)
 - [HealthInfo](docs/HealthInfo.md)
 - [IDReusePolicy](docs/IDReusePolicy.md)
 - [InterStateChannelCommand](docs/InterStateChannelCommand.md)
 - [InterStateChannelPublishing](docs/InterStateChannelPublishing.md)
 - [InterStateChannelResult](docs/InterStateChannelResult.md)
 - [KeyValue](docs/KeyValue.md)
 - [PersistenceLoadingPolicy](docs/PersistenceLoadingPolicy.md)
 - [PersistenceLoadingType](docs/PersistenceLoadingType.md)
 - [RetryPolicy](docs/RetryPolicy.md)
 - [SearchAttribute](docs/SearchAttribute.md)
 - [SearchAttributeKeyAndType](docs/SearchAttributeKeyAndType.md)
 - [SearchAttributeValueType](docs/SearchAttributeValueType.md)
 - [SignalCommand](docs/SignalCommand.md)
 - [SignalResult](docs/SignalResult.md)
 - [StateCompletionOutput](docs/StateCompletionOutput.md)
 - [StateDecision](docs/StateDecision.md)
 - [StateMovement](docs/StateMovement.md)
 - [TimerCommand](docs/TimerCommand.md)
 - [TimerResult](docs/TimerResult.md)
 - [TimerStatus](docs/TimerStatus.md)
 - [WaitUntilApiFailurePolicy](docs/WaitUntilApiFailurePolicy.md)
 - [WorkerErrorResponse](docs/WorkerErrorResponse.md)
 - [WorkflowConditionalClose](docs/WorkflowConditionalClose.md)
 - [WorkflowConditionalCloseType](docs/WorkflowConditionalCloseType.md)
 - [WorkflowConfig](docs/WorkflowConfig.md)
 - [WorkflowConfigUpdateRequest](docs/WorkflowConfigUpdateRequest.md)
 - [WorkflowDumpRequest](docs/WorkflowDumpRequest.md)
 - [WorkflowDumpResponse](docs/WorkflowDumpResponse.md)
 - [WorkflowErrorType](docs/WorkflowErrorType.md)
 - [WorkflowGetDataObjectsRequest](docs/WorkflowGetDataObjectsRequest.md)
 - [WorkflowGetDataObjectsResponse](docs/WorkflowGetDataObjectsResponse.md)
 - [WorkflowGetRequest](docs/WorkflowGetRequest.md)
 - [WorkflowGetResponse](docs/WorkflowGetResponse.md)
 - [WorkflowGetSearchAttributesRequest](docs/WorkflowGetSearchAttributesRequest.md)
 - [WorkflowGetSearchAttributesResponse](docs/WorkflowGetSearchAttributesResponse.md)
 - [WorkflowResetRequest](docs/WorkflowResetRequest.md)
 - [WorkflowResetResponse](docs/WorkflowResetResponse.md)
 - [WorkflowResetType](docs/WorkflowResetType.md)
 - [WorkflowRetryPolicy](docs/WorkflowRetryPolicy.md)
 - [WorkflowRpcRequest](docs/WorkflowRpcRequest.md)
 - [WorkflowRpcResponse](docs/WorkflowRpcResponse.md)
 - [WorkflowSearchRequest](docs/WorkflowSearchRequest.md)
 - [WorkflowSearchResponse](docs/WorkflowSearchResponse.md)
 - [WorkflowSearchResponseEntry](docs/WorkflowSearchResponseEntry.md)
 - [WorkflowSignalRequest](docs/WorkflowSignalRequest.md)
 - [WorkflowSkipTimerRequest](docs/WorkflowSkipTimerRequest.md)
 - [WorkflowStartOptions](docs/WorkflowStartOptions.md)
 - [WorkflowStartRequest](docs/WorkflowStartRequest.md)
 - [WorkflowStartResponse](docs/WorkflowStartResponse.md)
 - [WorkflowStateExecuteRequest](docs/WorkflowStateExecuteRequest.md)
 - [WorkflowStateExecuteResponse](docs/WorkflowStateExecuteResponse.md)
 - [WorkflowStateOptions](docs/WorkflowStateOptions.md)
 - [WorkflowStateWaitUntilRequest](docs/WorkflowStateWaitUntilRequest.md)
 - [WorkflowStateWaitUntilResponse](docs/WorkflowStateWaitUntilResponse.md)
 - [WorkflowStatus](docs/WorkflowStatus.md)
 - [WorkflowStopRequest](docs/WorkflowStopRequest.md)
 - [WorkflowStopType](docs/WorkflowStopType.md)
 - [WorkflowWaitForStateCompletionRequest](docs/WorkflowWaitForStateCompletionRequest.md)
 - [WorkflowWaitForStateCompletionResponse](docs/WorkflowWaitForStateCompletionResponse.md)
 - [WorkflowWorkerRpcRequest](docs/WorkflowWorkerRpcRequest.md)
 - [WorkflowWorkerRpcResponse](docs/WorkflowWorkerRpcResponse.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



