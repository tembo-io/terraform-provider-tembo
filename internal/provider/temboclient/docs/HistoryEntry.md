# HistoryEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**InstanceState**](InstanceState.md) |  | 
**Timestamp** | **time.Time** | The timestamp of the entry | 

## Methods

### NewHistoryEntry

`func NewHistoryEntry(instance InstanceState, timestamp time.Time, ) *HistoryEntry`

NewHistoryEntry instantiates a new HistoryEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryEntryWithDefaults

`func NewHistoryEntryWithDefaults() *HistoryEntry`

NewHistoryEntryWithDefaults instantiates a new HistoryEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *HistoryEntry) GetInstance() InstanceState`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *HistoryEntry) GetInstanceOk() (*InstanceState, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *HistoryEntry) SetInstance(v InstanceState)`

SetInstance sets Instance field to given value.


### GetTimestamp

`func (o *HistoryEntry) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *HistoryEntry) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *HistoryEntry) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


