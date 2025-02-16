# SearchParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **NullableInt32** | The number of results to return. Default is 10, max is 20. | [optional] 
**Query** | **string** | The search query. For example, \&quot;how to create a Tembo instance\&quot; | 

## Methods

### NewSearchParams

`func NewSearchParams(query string, ) *SearchParams`

NewSearchParams instantiates a new SearchParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchParamsWithDefaults

`func NewSearchParamsWithDefaults() *SearchParams`

NewSearchParamsWithDefaults instantiates a new SearchParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *SearchParams) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchParams) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchParams) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SearchParams) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *SearchParams) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *SearchParams) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetQuery

`func (o *SearchParams) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchParams) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchParams) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


