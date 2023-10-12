# WorkflowStateOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchAttributesLoadingPolicy** | Pointer to [**PersistenceLoadingPolicy**](PersistenceLoadingPolicy.md) |  | [optional] 
**DataAttributesLoadingPolicy** | Pointer to [**PersistenceLoadingPolicy**](PersistenceLoadingPolicy.md) |  | [optional] 
**WaitUntilApiTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**ExecuteApiTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**WaitUntilApiRetryPolicy** | Pointer to [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**ExecuteApiRetryPolicy** | Pointer to [**RetryPolicy**](RetryPolicy.md) |  | [optional] 
**WaitUntilApiFailurePolicy** | Pointer to [**WaitUntilApiFailurePolicy**](WaitUntilApiFailurePolicy.md) |  | [optional] 
**ExecuteApiFailurePolicy** | Pointer to [**ExecuteApiFailurePolicy**](ExecuteApiFailurePolicy.md) |  | [optional] 
**ExecuteApiFailureProceedStateId** | Pointer to **string** |  | [optional] 
**ExecuteApiFailureProceedStateOptions** | Pointer to [**WorkflowStateOptions**](WorkflowStateOptions.md) |  | [optional] 
**SkipWaitUntil** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkflowStateOptions

`func NewWorkflowStateOptions() *WorkflowStateOptions`

NewWorkflowStateOptions instantiates a new WorkflowStateOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStateOptionsWithDefaults

`func NewWorkflowStateOptionsWithDefaults() *WorkflowStateOptions`

NewWorkflowStateOptionsWithDefaults instantiates a new WorkflowStateOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchAttributesLoadingPolicy

`func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicy() PersistenceLoadingPolicy`

GetSearchAttributesLoadingPolicy returns the SearchAttributesLoadingPolicy field if non-nil, zero value otherwise.

### GetSearchAttributesLoadingPolicyOk

`func (o *WorkflowStateOptions) GetSearchAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool)`

GetSearchAttributesLoadingPolicyOk returns a tuple with the SearchAttributesLoadingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributesLoadingPolicy

`func (o *WorkflowStateOptions) SetSearchAttributesLoadingPolicy(v PersistenceLoadingPolicy)`

SetSearchAttributesLoadingPolicy sets SearchAttributesLoadingPolicy field to given value.

### HasSearchAttributesLoadingPolicy

`func (o *WorkflowStateOptions) HasSearchAttributesLoadingPolicy() bool`

HasSearchAttributesLoadingPolicy returns a boolean if a field has been set.

### GetDataAttributesLoadingPolicy

`func (o *WorkflowStateOptions) GetDataAttributesLoadingPolicy() PersistenceLoadingPolicy`

GetDataAttributesLoadingPolicy returns the DataAttributesLoadingPolicy field if non-nil, zero value otherwise.

### GetDataAttributesLoadingPolicyOk

`func (o *WorkflowStateOptions) GetDataAttributesLoadingPolicyOk() (*PersistenceLoadingPolicy, bool)`

GetDataAttributesLoadingPolicyOk returns a tuple with the DataAttributesLoadingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAttributesLoadingPolicy

`func (o *WorkflowStateOptions) SetDataAttributesLoadingPolicy(v PersistenceLoadingPolicy)`

SetDataAttributesLoadingPolicy sets DataAttributesLoadingPolicy field to given value.

### HasDataAttributesLoadingPolicy

`func (o *WorkflowStateOptions) HasDataAttributesLoadingPolicy() bool`

HasDataAttributesLoadingPolicy returns a boolean if a field has been set.

### GetWaitUntilApiTimeoutSeconds

`func (o *WorkflowStateOptions) GetWaitUntilApiTimeoutSeconds() int32`

GetWaitUntilApiTimeoutSeconds returns the WaitUntilApiTimeoutSeconds field if non-nil, zero value otherwise.

### GetWaitUntilApiTimeoutSecondsOk

`func (o *WorkflowStateOptions) GetWaitUntilApiTimeoutSecondsOk() (*int32, bool)`

GetWaitUntilApiTimeoutSecondsOk returns a tuple with the WaitUntilApiTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntilApiTimeoutSeconds

`func (o *WorkflowStateOptions) SetWaitUntilApiTimeoutSeconds(v int32)`

SetWaitUntilApiTimeoutSeconds sets WaitUntilApiTimeoutSeconds field to given value.

### HasWaitUntilApiTimeoutSeconds

`func (o *WorkflowStateOptions) HasWaitUntilApiTimeoutSeconds() bool`

HasWaitUntilApiTimeoutSeconds returns a boolean if a field has been set.

### GetExecuteApiTimeoutSeconds

`func (o *WorkflowStateOptions) GetExecuteApiTimeoutSeconds() int32`

GetExecuteApiTimeoutSeconds returns the ExecuteApiTimeoutSeconds field if non-nil, zero value otherwise.

### GetExecuteApiTimeoutSecondsOk

`func (o *WorkflowStateOptions) GetExecuteApiTimeoutSecondsOk() (*int32, bool)`

GetExecuteApiTimeoutSecondsOk returns a tuple with the ExecuteApiTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiTimeoutSeconds

`func (o *WorkflowStateOptions) SetExecuteApiTimeoutSeconds(v int32)`

SetExecuteApiTimeoutSeconds sets ExecuteApiTimeoutSeconds field to given value.

### HasExecuteApiTimeoutSeconds

`func (o *WorkflowStateOptions) HasExecuteApiTimeoutSeconds() bool`

HasExecuteApiTimeoutSeconds returns a boolean if a field has been set.

### GetWaitUntilApiRetryPolicy

`func (o *WorkflowStateOptions) GetWaitUntilApiRetryPolicy() RetryPolicy`

GetWaitUntilApiRetryPolicy returns the WaitUntilApiRetryPolicy field if non-nil, zero value otherwise.

### GetWaitUntilApiRetryPolicyOk

`func (o *WorkflowStateOptions) GetWaitUntilApiRetryPolicyOk() (*RetryPolicy, bool)`

GetWaitUntilApiRetryPolicyOk returns a tuple with the WaitUntilApiRetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntilApiRetryPolicy

`func (o *WorkflowStateOptions) SetWaitUntilApiRetryPolicy(v RetryPolicy)`

SetWaitUntilApiRetryPolicy sets WaitUntilApiRetryPolicy field to given value.

### HasWaitUntilApiRetryPolicy

`func (o *WorkflowStateOptions) HasWaitUntilApiRetryPolicy() bool`

HasWaitUntilApiRetryPolicy returns a boolean if a field has been set.

### GetExecuteApiRetryPolicy

`func (o *WorkflowStateOptions) GetExecuteApiRetryPolicy() RetryPolicy`

GetExecuteApiRetryPolicy returns the ExecuteApiRetryPolicy field if non-nil, zero value otherwise.

### GetExecuteApiRetryPolicyOk

`func (o *WorkflowStateOptions) GetExecuteApiRetryPolicyOk() (*RetryPolicy, bool)`

GetExecuteApiRetryPolicyOk returns a tuple with the ExecuteApiRetryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiRetryPolicy

`func (o *WorkflowStateOptions) SetExecuteApiRetryPolicy(v RetryPolicy)`

SetExecuteApiRetryPolicy sets ExecuteApiRetryPolicy field to given value.

### HasExecuteApiRetryPolicy

`func (o *WorkflowStateOptions) HasExecuteApiRetryPolicy() bool`

HasExecuteApiRetryPolicy returns a boolean if a field has been set.

### GetWaitUntilApiFailurePolicy

`func (o *WorkflowStateOptions) GetWaitUntilApiFailurePolicy() WaitUntilApiFailurePolicy`

GetWaitUntilApiFailurePolicy returns the WaitUntilApiFailurePolicy field if non-nil, zero value otherwise.

### GetWaitUntilApiFailurePolicyOk

`func (o *WorkflowStateOptions) GetWaitUntilApiFailurePolicyOk() (*WaitUntilApiFailurePolicy, bool)`

GetWaitUntilApiFailurePolicyOk returns a tuple with the WaitUntilApiFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntilApiFailurePolicy

`func (o *WorkflowStateOptions) SetWaitUntilApiFailurePolicy(v WaitUntilApiFailurePolicy)`

SetWaitUntilApiFailurePolicy sets WaitUntilApiFailurePolicy field to given value.

### HasWaitUntilApiFailurePolicy

`func (o *WorkflowStateOptions) HasWaitUntilApiFailurePolicy() bool`

HasWaitUntilApiFailurePolicy returns a boolean if a field has been set.

### GetExecuteApiFailurePolicy

`func (o *WorkflowStateOptions) GetExecuteApiFailurePolicy() ExecuteApiFailurePolicy`

GetExecuteApiFailurePolicy returns the ExecuteApiFailurePolicy field if non-nil, zero value otherwise.

### GetExecuteApiFailurePolicyOk

`func (o *WorkflowStateOptions) GetExecuteApiFailurePolicyOk() (*ExecuteApiFailurePolicy, bool)`

GetExecuteApiFailurePolicyOk returns a tuple with the ExecuteApiFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiFailurePolicy

`func (o *WorkflowStateOptions) SetExecuteApiFailurePolicy(v ExecuteApiFailurePolicy)`

SetExecuteApiFailurePolicy sets ExecuteApiFailurePolicy field to given value.

### HasExecuteApiFailurePolicy

`func (o *WorkflowStateOptions) HasExecuteApiFailurePolicy() bool`

HasExecuteApiFailurePolicy returns a boolean if a field has been set.

### GetExecuteApiFailureProceedStateId

`func (o *WorkflowStateOptions) GetExecuteApiFailureProceedStateId() string`

GetExecuteApiFailureProceedStateId returns the ExecuteApiFailureProceedStateId field if non-nil, zero value otherwise.

### GetExecuteApiFailureProceedStateIdOk

`func (o *WorkflowStateOptions) GetExecuteApiFailureProceedStateIdOk() (*string, bool)`

GetExecuteApiFailureProceedStateIdOk returns a tuple with the ExecuteApiFailureProceedStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiFailureProceedStateId

`func (o *WorkflowStateOptions) SetExecuteApiFailureProceedStateId(v string)`

SetExecuteApiFailureProceedStateId sets ExecuteApiFailureProceedStateId field to given value.

### HasExecuteApiFailureProceedStateId

`func (o *WorkflowStateOptions) HasExecuteApiFailureProceedStateId() bool`

HasExecuteApiFailureProceedStateId returns a boolean if a field has been set.

### GetExecuteApiFailureProceedStateOptions

`func (o *WorkflowStateOptions) GetExecuteApiFailureProceedStateOptions() WorkflowStateOptions`

GetExecuteApiFailureProceedStateOptions returns the ExecuteApiFailureProceedStateOptions field if non-nil, zero value otherwise.

### GetExecuteApiFailureProceedStateOptionsOk

`func (o *WorkflowStateOptions) GetExecuteApiFailureProceedStateOptionsOk() (*WorkflowStateOptions, bool)`

GetExecuteApiFailureProceedStateOptionsOk returns a tuple with the ExecuteApiFailureProceedStateOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteApiFailureProceedStateOptions

`func (o *WorkflowStateOptions) SetExecuteApiFailureProceedStateOptions(v WorkflowStateOptions)`

SetExecuteApiFailureProceedStateOptions sets ExecuteApiFailureProceedStateOptions field to given value.

### HasExecuteApiFailureProceedStateOptions

`func (o *WorkflowStateOptions) HasExecuteApiFailureProceedStateOptions() bool`

HasExecuteApiFailureProceedStateOptions returns a boolean if a field has been set.

### GetSkipWaitUntil

`func (o *WorkflowStateOptions) GetSkipWaitUntil() bool`

GetSkipWaitUntil returns the SkipWaitUntil field if non-nil, zero value otherwise.

### GetSkipWaitUntilOk

`func (o *WorkflowStateOptions) GetSkipWaitUntilOk() (*bool, bool)`

GetSkipWaitUntilOk returns a tuple with the SkipWaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipWaitUntil

`func (o *WorkflowStateOptions) SetSkipWaitUntil(v bool)`

SetSkipWaitUntil sets SkipWaitUntil field to given value.

### HasSkipWaitUntil

`func (o *WorkflowStateOptions) HasSkipWaitUntil() bool`

HasSkipWaitUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


