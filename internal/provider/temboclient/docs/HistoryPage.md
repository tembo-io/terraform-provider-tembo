# HistoryPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**History** | [**[]HistoryEntry**](HistoryEntry.md) | List of historical results | 
**Pagination** | [**PaginationInfo**](PaginationInfo.md) |  | 

## Methods

### NewHistoryPage

`func NewHistoryPage(history []HistoryEntry, pagination PaginationInfo, ) *HistoryPage`

NewHistoryPage instantiates a new HistoryPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryPageWithDefaults

`func NewHistoryPageWithDefaults() *HistoryPage`

NewHistoryPageWithDefaults instantiates a new HistoryPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistory

`func (o *HistoryPage) GetHistory() []HistoryEntry`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *HistoryPage) GetHistoryOk() (*[]HistoryEntry, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *HistoryPage) SetHistory(v []HistoryEntry)`

SetHistory sets History field to given value.


### GetPagination

`func (o *HistoryPage) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *HistoryPage) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *HistoryPage) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


