# ConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Password** | **string** |  | 
**PoolerHost** | Pointer to **NullableString** |  | [optional] 
**Port** | **int32** |  | 
**User** | **string** |  | 

## Methods

### NewConnectionInfo

`func NewConnectionInfo(host string, password string, port int32, user string, ) *ConnectionInfo`

NewConnectionInfo instantiates a new ConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInfoWithDefaults

`func NewConnectionInfoWithDefaults() *ConnectionInfo`

NewConnectionInfoWithDefaults instantiates a new ConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ConnectionInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ConnectionInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ConnectionInfo) SetHost(v string)`

SetHost sets Host field to given value.


### GetPassword

`func (o *ConnectionInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ConnectionInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ConnectionInfo) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPoolerHost

`func (o *ConnectionInfo) GetPoolerHost() string`

GetPoolerHost returns the PoolerHost field if non-nil, zero value otherwise.

### GetPoolerHostOk

`func (o *ConnectionInfo) GetPoolerHostOk() (*string, bool)`

GetPoolerHostOk returns a tuple with the PoolerHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolerHost

`func (o *ConnectionInfo) SetPoolerHost(v string)`

SetPoolerHost sets PoolerHost field to given value.

### HasPoolerHost

`func (o *ConnectionInfo) HasPoolerHost() bool`

HasPoolerHost returns a boolean if a field has been set.

### SetPoolerHostNil

`func (o *ConnectionInfo) SetPoolerHostNil(b bool)`

 SetPoolerHostNil sets the value for PoolerHost to be an explicit nil

### UnsetPoolerHost
`func (o *ConnectionInfo) UnsetPoolerHost()`

UnsetPoolerHost ensures that no value is present for PoolerHost, not even an explicit nil
### GetPort

`func (o *ConnectionInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ConnectionInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ConnectionInfo) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUser

`func (o *ConnectionInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConnectionInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConnectionInfo) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


