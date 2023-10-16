# Routing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressPath** | Pointer to **NullableString** |  | [optional] 
**Middlewares** | Pointer to **[]string** |  | [optional] 
**Port** | **int32** |  | 

## Methods

### NewRouting

`func NewRouting(port int32, ) *Routing`

NewRouting instantiates a new Routing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingWithDefaults

`func NewRoutingWithDefaults() *Routing`

NewRoutingWithDefaults instantiates a new Routing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressPath

`func (o *Routing) GetIngressPath() string`

GetIngressPath returns the IngressPath field if non-nil, zero value otherwise.

### GetIngressPathOk

`func (o *Routing) GetIngressPathOk() (*string, bool)`

GetIngressPathOk returns a tuple with the IngressPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPath

`func (o *Routing) SetIngressPath(v string)`

SetIngressPath sets IngressPath field to given value.

### HasIngressPath

`func (o *Routing) HasIngressPath() bool`

HasIngressPath returns a boolean if a field has been set.

### SetIngressPathNil

`func (o *Routing) SetIngressPathNil(b bool)`

 SetIngressPathNil sets the value for IngressPath to be an explicit nil

### UnsetIngressPath
`func (o *Routing) UnsetIngressPath()`

UnsetIngressPath ensures that no value is present for IngressPath, not even an explicit nil
### GetMiddlewares

`func (o *Routing) GetMiddlewares() []string`

GetMiddlewares returns the Middlewares field if non-nil, zero value otherwise.

### GetMiddlewaresOk

`func (o *Routing) GetMiddlewaresOk() (*[]string, bool)`

GetMiddlewaresOk returns a tuple with the Middlewares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddlewares

`func (o *Routing) SetMiddlewares(v []string)`

SetMiddlewares sets Middlewares field to given value.

### HasMiddlewares

`func (o *Routing) HasMiddlewares() bool`

HasMiddlewares returns a boolean if a field has been set.

### SetMiddlewaresNil

`func (o *Routing) SetMiddlewaresNil(b bool)`

 SetMiddlewaresNil sets the value for Middlewares to be an explicit nil

### UnsetMiddlewares
`func (o *Routing) UnsetMiddlewares()`

UnsetMiddlewares ensures that no value is present for Middlewares, not even an explicit nil
### GetPort

`func (o *Routing) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Routing) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Routing) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


