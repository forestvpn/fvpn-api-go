/*
ForestVPN API

Testing DeviceAPIService

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

func Test_fvpn_DeviceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeviceAPIService CreateDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeviceAPI.CreateDevice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService DeleteDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		httpRes, err := apiClient.DeviceAPI.DeleteDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService GetDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceAPI.GetDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService GetDevices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeviceAPI.GetDevices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService PartialUpdateDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceAPI.PartialUpdateDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeviceAPIService UpdateDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var deviceId string

		resp, httpRes, err := apiClient.DeviceAPI.UpdateDevice(context.Background(), deviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
