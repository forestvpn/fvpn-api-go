# SetNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The ID of the node to set | [optional] 
**Lat** | Pointer to **float64** | latitude coordinate for location-based node selection | [optional] 
**Lon** | Pointer to **float64** | longitude coordinate for location-based node selection | [optional] 
**MaxDistance** | Pointer to **string** | Maximum distance from coordinates (format: &#39;10km&#39;, &#39;5mi&#39;) | [optional] [default to "500km"]

## Methods

### NewSetNodeRequest

`func NewSetNodeRequest() *SetNodeRequest`

NewSetNodeRequest instantiates a new SetNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodeRequestWithDefaults

`func NewSetNodeRequestWithDefaults() *SetNodeRequest`

NewSetNodeRequestWithDefaults instantiates a new SetNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *SetNodeRequest) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SetNodeRequest) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SetNodeRequest) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *SetNodeRequest) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetLat

`func (o *SetNodeRequest) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *SetNodeRequest) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *SetNodeRequest) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *SetNodeRequest) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *SetNodeRequest) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *SetNodeRequest) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *SetNodeRequest) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *SetNodeRequest) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetMaxDistance

`func (o *SetNodeRequest) GetMaxDistance() string`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *SetNodeRequest) GetMaxDistanceOk() (*string, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *SetNodeRequest) SetMaxDistance(v string)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *SetNodeRequest) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


