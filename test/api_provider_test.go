/*
ForestVPN API

Testing ProviderAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fvpn

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/forestvpn/fvpn-api-go"
)

func Test_fvpn_ProviderAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProviderAPIService GetProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.ProviderAPI.GetProvider(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProviderAPIService GetProviders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProviderAPI.GetProviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
