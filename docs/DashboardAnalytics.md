# DashboardAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary7d** | [**ComparisonData**](ComparisonData.md) |  | 
**Summary30d** | [**ComparisonData**](ComparisonData.md) |  | 
**TopDevices** | [**TopDevices**](TopDevices.md) |  | 

## Methods

### NewDashboardAnalytics

`func NewDashboardAnalytics(summary7d ComparisonData, summary30d ComparisonData, topDevices TopDevices, ) *DashboardAnalytics`

NewDashboardAnalytics instantiates a new DashboardAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardAnalyticsWithDefaults

`func NewDashboardAnalyticsWithDefaults() *DashboardAnalytics`

NewDashboardAnalyticsWithDefaults instantiates a new DashboardAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary7d

`func (o *DashboardAnalytics) GetSummary7d() ComparisonData`

GetSummary7d returns the Summary7d field if non-nil, zero value otherwise.

### GetSummary7dOk

`func (o *DashboardAnalytics) GetSummary7dOk() (*ComparisonData, bool)`

GetSummary7dOk returns a tuple with the Summary7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary7d

`func (o *DashboardAnalytics) SetSummary7d(v ComparisonData)`

SetSummary7d sets Summary7d field to given value.


### GetSummary30d

`func (o *DashboardAnalytics) GetSummary30d() ComparisonData`

GetSummary30d returns the Summary30d field if non-nil, zero value otherwise.

### GetSummary30dOk

`func (o *DashboardAnalytics) GetSummary30dOk() (*ComparisonData, bool)`

GetSummary30dOk returns a tuple with the Summary30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary30d

`func (o *DashboardAnalytics) SetSummary30d(v ComparisonData)`

SetSummary30d sets Summary30d field to given value.


### GetTopDevices

`func (o *DashboardAnalytics) GetTopDevices() TopDevices`

GetTopDevices returns the TopDevices field if non-nil, zero value otherwise.

### GetTopDevicesOk

`func (o *DashboardAnalytics) GetTopDevicesOk() (*TopDevices, bool)`

GetTopDevicesOk returns a tuple with the TopDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopDevices

`func (o *DashboardAnalytics) SetTopDevices(v TopDevices)`

SetTopDevices sets TopDevices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


