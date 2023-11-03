/*
Tembo Cloud

Testing InstanceApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package temboclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_temboclient_InstanceApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InstanceApiService CreateInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.InstanceApi.CreateInstance(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService DeleteInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var instanceId string

		resp, httpRes, err := apiClient.InstanceApi.DeleteInstance(context.Background(), orgId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService GetAll", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.InstanceApi.GetAll(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService GetInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var instanceId string

		resp, httpRes, err := apiClient.InstanceApi.GetInstance(context.Background(), orgId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService GetSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstanceApi.GetSchema(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService InstanceEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var instanceId string

		resp, httpRes, err := apiClient.InstanceApi.InstanceEvent(context.Background(), orgId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService PatchInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var instanceId string

		resp, httpRes, err := apiClient.InstanceApi.PatchInstance(context.Background(), orgId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService PutInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var instanceId string

		resp, httpRes, err := apiClient.InstanceApi.PutInstance(context.Background(), orgId, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstanceApiService RestoreInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.InstanceApi.RestoreInstance(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
