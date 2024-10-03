# TimerCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandId** | Pointer to **string** |  | [optional] 
**DurationSeconds** | **int64** |  | 

## Methods

### NewTimerCommand

`func NewTimerCommand(durationSeconds int64, ) *TimerCommand`

NewTimerCommand instantiates a new TimerCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerCommandWithDefaults

`func NewTimerCommandWithDefaults() *TimerCommand`

NewTimerCommandWithDefaults instantiates a new TimerCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandId

`func (o *TimerCommand) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *TimerCommand) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *TimerCommand) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *TimerCommand) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *TimerCommand) GetDurationSeconds() int64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *TimerCommand) GetDurationSecondsOk() (*int64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *TimerCommand) SetDurationSeconds(v int64)`

SetDurationSeconds sets DurationSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


