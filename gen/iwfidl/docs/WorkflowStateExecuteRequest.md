# WorkflowStateExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**WorkflowType** | **string** |  | 
**WorkflowStateId** | **string** |  | 
**StateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**SearchAttributes** | Pointer to [**[]SearchAttribute**](SearchAttribute.md) |  | [optional] 
**DataObjects** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**StateLocals** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**CommandResults** | Pointer to [**CommandResults**](CommandResults.md) |  | [optional] 

## Methods

### NewWorkflowStateExecuteRequest

`func NewWorkflowStateExecuteRequest(context Context, workflowType string, workflowStateId string, ) *WorkflowStateExecuteRequest`

NewWorkflowStateExecuteRequest instantiates a new WorkflowStateExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateExecuteRequestWithDefaults

`func NewWorkflowStateExecuteRequestWithDefaults() *WorkflowStateExecuteRequest`

NewWorkflowStateExecuteRequestWithDefaults instantiates a new WorkflowStateExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *WorkflowStateExecuteRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WorkflowStateExecuteRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WorkflowStateExecuteRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetWorkflowType

`func (o *WorkflowStateExecuteRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowStateExecuteRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowStateExecuteRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### GetWorkflowStateId

`func (o *WorkflowStateExecuteRequest) GetWorkflowStateId() string`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *WorkflowStateExecuteRequest) GetWorkflowStateIdOk() (*string, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *WorkflowStateExecuteRequest) SetWorkflowStateId(v string)`

SetWorkflowStateId sets WorkflowStateId field to given value.


### GetStateInput

`func (o *WorkflowStateExecuteRequest) GetStateInput() EncodedObject`

GetStateInput returns the StateInput field if non-nil, zero value otherwise.

### GetStateInputOk

`func (o *WorkflowStateExecuteRequest) GetStateInputOk() (*EncodedObject, bool)`

GetStateInputOk returns a tuple with the StateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInput

`func (o *WorkflowStateExecuteRequest) SetStateInput(v EncodedObject)`

SetStateInput sets StateInput field to given value.

### HasStateInput

`func (o *WorkflowStateExecuteRequest) HasStateInput() bool`

HasStateInput returns a boolean if a field has been set.

### GetSearchAttributes

`func (o *WorkflowStateExecuteRequest) GetSearchAttributes() []SearchAttribute`

GetSearchAttributes returns the SearchAttributes field if non-nil, zero value otherwise.

### GetSearchAttributesOk

`func (o *WorkflowStateExecuteRequest) GetSearchAttributesOk() (*[]SearchAttribute, bool)`

GetSearchAttributesOk returns a tuple with the SearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributes

`func (o *WorkflowStateExecuteRequest) SetSearchAttributes(v []SearchAttribute)`

SetSearchAttributes sets SearchAttributes field to given value.

### HasSearchAttributes

`func (o *WorkflowStateExecuteRequest) HasSearchAttributes() bool`

HasSearchAttributes returns a boolean if a field has been set.

### GetDataObjects

`func (o *WorkflowStateExecuteRequest) GetDataObjects() []KeyValue`

GetDataObjects returns the DataObjects field if non-nil, zero value otherwise.

### GetDataObjectsOk

`func (o *WorkflowStateExecuteRequest) GetDataObjectsOk() (*[]KeyValue, bool)`

GetDataObjectsOk returns a tuple with the DataObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataObjects

`func (o *WorkflowStateExecuteRequest) SetDataObjects(v []KeyValue)`

SetDataObjects sets DataObjects field to given value.

### HasDataObjects

`func (o *WorkflowStateExecuteRequest) HasDataObjects() bool`

HasDataObjects returns a boolean if a field has been set.

### GetStateLocals

`func (o *WorkflowStateExecuteRequest) GetStateLocals() []KeyValue`

GetStateLocals returns the StateLocals field if non-nil, zero value otherwise.

### GetStateLocalsOk

`func (o *WorkflowStateExecuteRequest) GetStateLocalsOk() (*[]KeyValue, bool)`

GetStateLocalsOk returns a tuple with the StateLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateLocals

`func (o *WorkflowStateExecuteRequest) SetStateLocals(v []KeyValue)`

SetStateLocals sets StateLocals field to given value.

### HasStateLocals

`func (o *WorkflowStateExecuteRequest) HasStateLocals() bool`

HasStateLocals returns a boolean if a field has been set.

### GetCommandResults

`func (o *WorkflowStateExecuteRequest) GetCommandResults() CommandResults`

GetCommandResults returns the CommandResults field if non-nil, zero value otherwise.

### GetCommandResultsOk

`func (o *WorkflowStateExecuteRequest) GetCommandResultsOk() (*CommandResults, bool)`

GetCommandResultsOk returns a tuple with the CommandResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandResults

`func (o *WorkflowStateExecuteRequest) SetCommandResults(v CommandResults)`

SetCommandResults sets CommandResults field to given value.

### HasCommandResults

`func (o *WorkflowStateExecuteRequest) HasCommandResults() bool`

HasCommandResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


