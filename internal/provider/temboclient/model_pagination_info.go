/*
Tembo Cloud

Platform API for Tembo Cloud             </br>             </br>             To find a Tembo Data API, please find it here:             </br>             </br>             [AWS US East 1](https://api.data-1.use1.tembo.io/swagger-ui/)             

API version: v1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package temboclient

import (
	"encoding/json"
	"fmt"
)

// checks if the PaginationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationInfo{}

// PaginationInfo struct for PaginationInfo
type PaginationInfo struct {
	// The page returned
	CurrentPage int32 `json:"currentPage"`
	// True if more results are available on the next page
	HasMore bool `json:"hasMore"`
	// The maximum number of results returned
	Limit int32 `json:"limit"`
	AdditionalProperties map[string]interface{}
}

type _PaginationInfo PaginationInfo

// NewPaginationInfo instantiates a new PaginationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationInfo(currentPage int32, hasMore bool, limit int32) *PaginationInfo {
	this := PaginationInfo{}
	this.CurrentPage = currentPage
	this.HasMore = hasMore
	this.Limit = limit
	return &this
}

// NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationInfoWithDefaults() *PaginationInfo {
	this := PaginationInfo{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value
func (o *PaginationInfo) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *PaginationInfo) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetHasMore returns the HasMore field value
func (o *PaginationInfo) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value
func (o *PaginationInfo) SetHasMore(v bool) {
	o.HasMore = v
}

// GetLimit returns the Limit field value
func (o *PaginationInfo) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PaginationInfo) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *PaginationInfo) SetLimit(v int32) {
	o.Limit = v
}

func (o PaginationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentPage"] = o.CurrentPage
	toSerialize["hasMore"] = o.HasMore
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginationInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentPage",
		"hasMore",
		"limit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPaginationInfo := _PaginationInfo{}

	err = json.Unmarshal(data, &varPaginationInfo)

	if err != nil {
		return err
	}

	*o = PaginationInfo(varPaginationInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "currentPage")
		delete(additionalProperties, "hasMore")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginationInfo struct {
	value *PaginationInfo
	isSet bool
}

func (v NullablePaginationInfo) Get() *PaginationInfo {
	return v.value
}

func (v *NullablePaginationInfo) Set(val *PaginationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationInfo(val *PaginationInfo) *NullablePaginationInfo {
	return &NullablePaginationInfo{value: val, isSet: true}
}

func (v NullablePaginationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


