# WorkflowStateWaitUntilRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**Context**](Context.md) |  | 
**WorkflowType** | **string** |  | 
**WorkflowStateId** | **string** |  | 
**StateInput** | Pointer to [**EncodedObject**](EncodedObject.md) |  | [optional] 
**SearchAttributes** | Pointer to [**[]SearchAttribute**](SearchAttribute.md) |  | [optional] 
**DataObjects** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 

## Methods

### NewWorkflowStateWaitUntilRequest

`func NewWorkflowStateWaitUntilRequest(context Context, workflowType string, workflowStateId string, ) *WorkflowStateWaitUntilRequest`

NewWorkflowStateWaitUntilRequest instantiates a new WorkflowStateWaitUntilRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateWaitUntilRequestWithDefaults

`func NewWorkflowStateWaitUntilRequestWithDefaults() *WorkflowStateWaitUntilRequest`

NewWorkflowStateWaitUntilRequestWithDefaults instantiates a new WorkflowStateWaitUntilRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *WorkflowStateWaitUntilRequest) GetContext() Context`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *WorkflowStateWaitUntilRequest) GetContextOk() (*Context, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *WorkflowStateWaitUntilRequest) SetContext(v Context)`

SetContext sets Context field to given value.


### GetWorkflowType

`func (o *WorkflowStateWaitUntilRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowStateWaitUntilRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowStateWaitUntilRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### GetWorkflowStateId

`func (o *WorkflowStateWaitUntilRequest) GetWorkflowStateId() string`

GetWorkflowStateId returns the WorkflowStateId field if non-nil, zero value otherwise.

### GetWorkflowStateIdOk

`func (o *WorkflowStateWaitUntilRequest) GetWorkflowStateIdOk() (*string, bool)`

GetWorkflowStateIdOk returns a tuple with the WorkflowStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStateId

`func (o *WorkflowStateWaitUntilRequest) SetWorkflowStateId(v string)`

SetWorkflowStateId sets WorkflowStateId field to given value.


### GetStateInput

`func (o *WorkflowStateWaitUntilRequest) GetStateInput() EncodedObject`

GetStateInput returns the StateInput field if non-nil, zero value otherwise.

### GetStateInputOk

`func (o *WorkflowStateWaitUntilRequest) GetStateInputOk() (*EncodedObject, bool)`

GetStateInputOk returns a tuple with the StateInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInput

`func (o *WorkflowStateWaitUntilRequest) SetStateInput(v EncodedObject)`

SetStateInput sets StateInput field to given value.

### HasStateInput

`func (o *WorkflowStateWaitUntilRequest) HasStateInput() bool`

HasStateInput returns a boolean if a field has been set.

### GetSearchAttributes

`func (o *WorkflowStateWaitUntilRequest) GetSearchAttributes() []SearchAttribute`

GetSearchAttributes returns the SearchAttributes field if non-nil, zero value otherwise.

### GetSearchAttributesOk

`func (o *WorkflowStateWaitUntilRequest) GetSearchAttributesOk() (*[]SearchAttribute, bool)`

GetSearchAttributesOk returns a tuple with the SearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributes

`func (o *WorkflowStateWaitUntilRequest) SetSearchAttributes(v []SearchAttribute)`

SetSearchAttributes sets SearchAttributes field to given value.

### HasSearchAttributes

`func (o *WorkflowStateWaitUntilRequest) HasSearchAttributes() bool`

HasSearchAttributes returns a boolean if a field has been set.

### GetDataObjects

`func (o *WorkflowStateWaitUntilRequest) GetDataObjects() []KeyValue`

GetDataObjects returns the DataObjects field if non-nil, zero value otherwise.

### GetDataObjectsOk

`func (o *WorkflowStateWaitUntilRequest) GetDataObjectsOk() (*[]KeyValue, bool)`

GetDataObjectsOk returns a tuple with the DataObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataObjects

`func (o *WorkflowStateWaitUntilRequest) SetDataObjects(v []KeyValue)`

SetDataObjects sets DataObjects field to given value.

### HasDataObjects

`func (o *WorkflowStateWaitUntilRequest) HasDataObjects() bool`

HasDataObjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


