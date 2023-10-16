# Middleware

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomRequestHeaders** | [**HeaderConfig**](HeaderConfig.md) |  | 
**StripPrefix** | [**StripPrefixConfig**](StripPrefixConfig.md) |  | 

## Methods

### NewMiddleware

`func NewMiddleware(customRequestHeaders HeaderConfig, stripPrefix StripPrefixConfig, ) *Middleware`

NewMiddleware instantiates a new Middleware object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMiddlewareWithDefaults

`func NewMiddlewareWithDefaults() *Middleware`

NewMiddlewareWithDefaults instantiates a new Middleware object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomRequestHeaders

`func (o *Middleware) GetCustomRequestHeaders() HeaderConfig`

GetCustomRequestHeaders returns the CustomRequestHeaders field if non-nil, zero value otherwise.

### GetCustomRequestHeadersOk

`func (o *Middleware) GetCustomRequestHeadersOk() (*HeaderConfig, bool)`

GetCustomRequestHeadersOk returns a tuple with the CustomRequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRequestHeaders

`func (o *Middleware) SetCustomRequestHeaders(v HeaderConfig)`

SetCustomRequestHeaders sets CustomRequestHeaders field to given value.


### GetStripPrefix

`func (o *Middleware) GetStripPrefix() StripPrefixConfig`

GetStripPrefix returns the StripPrefix field if non-nil, zero value otherwise.

### GetStripPrefixOk

`func (o *Middleware) GetStripPrefixOk() (*StripPrefixConfig, bool)`

GetStripPrefixOk returns a tuple with the StripPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripPrefix

`func (o *Middleware) SetStripPrefix(v StripPrefixConfig)`

SetStripPrefix sets StripPrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


