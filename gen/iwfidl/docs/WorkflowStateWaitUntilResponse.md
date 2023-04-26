# WorkflowStateWaitUntilResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpsertSearchAttributes** | Pointer to [**[]SearchAttribute**](SearchAttribute.md) |  | [optional] 
**UpsertDataObjects** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**CommandRequest** | Pointer to [**CommandRequest**](CommandRequest.md) |  | [optional] 
**UpsertStateLocals** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**RecordEvents** | Pointer to [**[]KeyValue**](KeyValue.md) |  | [optional] 
**PublishToInterStateChannel** | Pointer to [**[]InterStateChannelPublishing**](InterStateChannelPublishing.md) |  | [optional] 

## Methods

### NewWorkflowStateWaitUntilResponse

`func NewWorkflowStateWaitUntilResponse() *WorkflowStateWaitUntilResponse`

NewWorkflowStateWaitUntilResponse instantiates a new WorkflowStateWaitUntilResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateWaitUntilResponseWithDefaults

`func NewWorkflowStateWaitUntilResponseWithDefaults() *WorkflowStateWaitUntilResponse`

NewWorkflowStateWaitUntilResponseWithDefaults instantiates a new WorkflowStateWaitUntilResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpsertSearchAttributes

`func (o *WorkflowStateWaitUntilResponse) GetUpsertSearchAttributes() []SearchAttribute`

GetUpsertSearchAttributes returns the UpsertSearchAttributes field if non-nil, zero value otherwise.

### GetUpsertSearchAttributesOk

`func (o *WorkflowStateWaitUntilResponse) GetUpsertSearchAttributesOk() (*[]SearchAttribute, bool)`

GetUpsertSearchAttributesOk returns a tuple with the UpsertSearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertSearchAttributes

`func (o *WorkflowStateWaitUntilResponse) SetUpsertSearchAttributes(v []SearchAttribute)`

SetUpsertSearchAttributes sets UpsertSearchAttributes field to given value.

### HasUpsertSearchAttributes

`func (o *WorkflowStateWaitUntilResponse) HasUpsertSearchAttributes() bool`

HasUpsertSearchAttributes returns a boolean if a field has been set.

### GetUpsertDataObjects

`func (o *WorkflowStateWaitUntilResponse) GetUpsertDataObjects() []KeyValue`

GetUpsertDataObjects returns the UpsertDataObjects field if non-nil, zero value otherwise.

### GetUpsertDataObjectsOk

`func (o *WorkflowStateWaitUntilResponse) GetUpsertDataObjectsOk() (*[]KeyValue, bool)`

GetUpsertDataObjectsOk returns a tuple with the UpsertDataObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertDataObjects

`func (o *WorkflowStateWaitUntilResponse) SetUpsertDataObjects(v []KeyValue)`

SetUpsertDataObjects sets UpsertDataObjects field to given value.

### HasUpsertDataObjects

`func (o *WorkflowStateWaitUntilResponse) HasUpsertDataObjects() bool`

HasUpsertDataObjects returns a boolean if a field has been set.

### GetCommandRequest

`func (o *WorkflowStateWaitUntilResponse) GetCommandRequest() CommandRequest`

GetCommandRequest returns the CommandRequest field if non-nil, zero value otherwise.

### GetCommandRequestOk

`func (o *WorkflowStateWaitUntilResponse) GetCommandRequestOk() (*CommandRequest, bool)`

GetCommandRequestOk returns a tuple with the CommandRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandRequest

`func (o *WorkflowStateWaitUntilResponse) SetCommandRequest(v CommandRequest)`

SetCommandRequest sets CommandRequest field to given value.

### HasCommandRequest

`func (o *WorkflowStateWaitUntilResponse) HasCommandRequest() bool`

HasCommandRequest returns a boolean if a field has been set.

### GetUpsertStateLocals

`func (o *WorkflowStateWaitUntilResponse) GetUpsertStateLocals() []KeyValue`

GetUpsertStateLocals returns the UpsertStateLocals field if non-nil, zero value otherwise.

### GetUpsertStateLocalsOk

`func (o *WorkflowStateWaitUntilResponse) GetUpsertStateLocalsOk() (*[]KeyValue, bool)`

GetUpsertStateLocalsOk returns a tuple with the UpsertStateLocals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpsertStateLocals

`func (o *WorkflowStateWaitUntilResponse) SetUpsertStateLocals(v []KeyValue)`

SetUpsertStateLocals sets UpsertStateLocals field to given value.

### HasUpsertStateLocals

`func (o *WorkflowStateWaitUntilResponse) HasUpsertStateLocals() bool`

HasUpsertStateLocals returns a boolean if a field has been set.

### GetRecordEvents

`func (o *WorkflowStateWaitUntilResponse) GetRecordEvents() []KeyValue`

GetRecordEvents returns the RecordEvents field if non-nil, zero value otherwise.

### GetRecordEventsOk

`func (o *WorkflowStateWaitUntilResponse) GetRecordEventsOk() (*[]KeyValue, bool)`

GetRecordEventsOk returns a tuple with the RecordEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordEvents

`func (o *WorkflowStateWaitUntilResponse) SetRecordEvents(v []KeyValue)`

SetRecordEvents sets RecordEvents field to given value.

### HasRecordEvents

`func (o *WorkflowStateWaitUntilResponse) HasRecordEvents() bool`

HasRecordEvents returns a boolean if a field has been set.

### GetPublishToInterStateChannel

`func (o *WorkflowStateWaitUntilResponse) GetPublishToInterStateChannel() []InterStateChannelPublishing`

GetPublishToInterStateChannel returns the PublishToInterStateChannel field if non-nil, zero value otherwise.

### GetPublishToInterStateChannelOk

`func (o *WorkflowStateWaitUntilResponse) GetPublishToInterStateChannelOk() (*[]InterStateChannelPublishing, bool)`

GetPublishToInterStateChannelOk returns a tuple with the PublishToInterStateChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishToInterStateChannel

`func (o *WorkflowStateWaitUntilResponse) SetPublishToInterStateChannel(v []InterStateChannelPublishing)`

SetPublishToInterStateChannel sets PublishToInterStateChannel field to given value.

### HasPublishToInterStateChannel

`func (o *WorkflowStateWaitUntilResponse) HasPublishToInterStateChannel() bool`

HasPublishToInterStateChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


