# WorkflowStateExecuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateDecision** | Pointer to [**StateDecision**](StateDecision.md) |  | [optional] 
**UpsertSearchAttributes** | Pointer to [**[]SearchAttribute**](SearchAttribute.md) |  | [optional] 
**UpsertDataObjects** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**RecordEvents** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**UpsertStateLocals** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**PublishToInterStateChannel** | Pointer to [**[]InterStateChannelPublishing**](InterStateChannelPublishing.md) |  | [optional] 

## Methods

### NewWorkflowStateExecuteResponse

`func NewWorkflowStateExecuteResponse() *WorkflowStateExecuteResponse`

NewWorkflowStateExecuteResponse instantiates a new WorkflowStateExecuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateExecuteResponseWithDefaults

`func NewWorkflowStateExecuteResponseWithDefaults() *WorkflowStateExecuteResponse`

NewWorkflowStateExecuteResponseWithDefaults instantiates a new WorkflowStateExecuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateDecision

`func (o *WorkflowStateExecuteResponse) GetStateDecision() StateDecision`

GetStateDecision returns the StateDecision field if non-nil, zero value otherwise.

### GetStateDecisionOk

`func (o *WorkflowStateExecuteResponse) GetStateDecisionOk() (*StateDecision, bool)`

GetStateDecisionOk returns a tuple with the StateDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDecision

`func (o *WorkflowStateExecuteResponse) SetStateDecision(v StateDecision)`

SetStateDecision sets StateDecision field to given value.

### HasStateDecision

`func (o *WorkflowStateExecuteResponse) HasStateDecision() bool`

HasStateDecision returns a boolean if a field has been set.

### GetUpsertSearchAttributes

`func (o *WorkflowStateExecuteResponse) GetUpsertSearchAttributes() []SearchAttribute`

GetUpsertSearchAttributes returns the UpsertSearchAttributes field if non-nil, zero value otherwise.

### GetUpsertSearchAttributesOk

`func (o *WorkflowStateExecuteResponse) GetUpsertSearchAttributesOk() (*[]SearchAttribute, bool)`

GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertSearchAttributes

`func (o *WorkflowStateExecuteResponse) SetUpsertSearchAttributes(v []SearchAttribute)`

SetUpsertSearchAttributes sets UpsertSearchAttributes field to given value.

### HasUpsertSearchAttributes

`func (o *WorkflowStateExecuteResponse) HasUpsertSearchAttributes() bool`

HasUpsertSearchAttributes returns a boolean if a field has been set.

### GetUpsertDataObjects

`func (o *WorkflowStateExecuteResponse) GetUpsertDataObjects() []KeyValue`

GetUpsertDataObjects returns the UpsertDataObjects field if non-nil, zero value otherwise.

### GetUpsertDataObjectsOk

`func (o *WorkflowStateExecuteResponse) GetUpsertDataObjectsOk() (*[]KeyValue, bool)`

GetUpsertDataObjectsOk returns a tuple with the UpsertDataObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertDataObjects

`func (o *WorkflowStateExecuteResponse) SetUpsertDataObjects(v []KeyValue)`

SetUpsertDataObjects sets UpsertDataObjects field to given value.

### HasUpsertDataObjects

`func (o *WorkflowStateExecuteResponse) HasUpsertDataObjects() bool`

HasUpsertDataObjects returns a boolean if a field has been set.

### GetRecordEvents

`func (o *WorkflowStateExecuteResponse) GetRecordEvents() []KeyValue`

GetRecordEvents returns the RecordEvents field if non-nil, zero value otherwise.

### GetRecordEventsOk

`func (o *WorkflowStateExecuteResponse) GetRecordEventsOk() (*[]KeyValue, bool)`

GetRecordEventsOk returns a tuple with the RecordEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordEvents

`func (o *WorkflowStateExecuteResponse) SetRecordEvents(v []KeyValue)`

SetRecordEvents sets RecordEvents field to given value.

### HasRecordEvents

`func (o *WorkflowStateExecuteResponse) HasRecordEvents() bool`

HasRecordEvents returns a boolean if a field has been set.

### GetUpsertStateLocals

`func (o *WorkflowStateExecuteResponse) GetUpsertStateLocals() []KeyValue`

GetUpsertStateLocals returns the UpsertStateLocals field if non-nil, zero value otherwise.

### GetUpsertStateLocalsOk

`func (o *WorkflowStateExecuteResponse) GetUpsertStateLocalsOk() (*[]KeyValue, bool)`

GetUpsertStateLocalsOk returns a tuple with the UpsertStateLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertStateLocals

`func (o *WorkflowStateExecuteResponse) SetUpsertStateLocals(v []KeyValue)`

SetUpsertStateLocals sets UpsertStateLocals field to given value.

### HasUpsertStateLocals

`func (o *WorkflowStateExecuteResponse) HasUpsertStateLocals() bool`

HasUpsertStateLocals returns a boolean if a field has been set.

### GetPublishToInterStateChannel

`func (o *WorkflowStateExecuteResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing`

GetPublishToInterStateChannel returns the PublishToInterStateChannel field if non-nil, zero value otherwise.

### GetPublishToInterStateChannelOk

`func (o *WorkflowStateExecuteResponse) GetPublishToInterStateChannelOk() (*[]InterStateChannelPublishing, bool)`

GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToInterStateChannel

`func (o *WorkflowStateExecuteResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing)`

SetPublishToInterStateChannel sets PublishToInterStateChannel field to given value.

### HasPublishToInterStateChannel

`func (o *WorkflowStateExecuteResponse) HasPublishToInterStateChannel() bool`

HasPublishToInterStateChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


