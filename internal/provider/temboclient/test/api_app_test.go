/*
Tembo Cloud

Testing AppAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package temboclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func Test_temboclient_AppAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AppAPIService GetAllApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AppAPI.GetAllApps(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AppAPIService GetApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.AppAPI.GetApp(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
