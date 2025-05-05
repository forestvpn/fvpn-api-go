# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Country** | [**Country**](Country.md) |  | 
**Region** | [**LocationRegion**](LocationRegion.md) |  | 
**Name** | **string** |  | 
**NameAscii** | Pointer to **NullableString** |  | [optional] 
**Lat** | Pointer to **NullableFloat64** |  | [optional] 
**Lon** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewLocation

`func NewLocation(id string, country Country, region LocationRegion, name string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v string)`

SetId sets Id field to given value.


### GetCountry

`func (o *Location) GetCountry() Country`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*Country, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v Country)`

SetCountry sets Country field to given value.


### GetRegion

`func (o *Location) GetRegion() LocationRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Location) GetRegionOk() (*LocationRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Location) SetRegion(v LocationRegion)`

SetRegion sets Region field to given value.


### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.


### GetNameAscii

`func (o *Location) GetNameAscii() string`

GetNameAscii returns the NameAscii field if non-nil, zero value otherwise.

### GetNameAsciiOk

`func (o *Location) GetNameAsciiOk() (*string, bool)`

GetNameAsciiOk returns a tuple with the NameAscii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameAscii

`func (o *Location) SetNameAscii(v string)`

SetNameAscii sets NameAscii field to given value.

### HasNameAscii

`func (o *Location) HasNameAscii() bool`

HasNameAscii returns a boolean if a field has been set.

### SetNameAsciiNil

`func (o *Location) SetNameAsciiNil(b bool)`

 SetNameAsciiNil sets the value for NameAscii to be an explicit nil

### UnsetNameAscii
`func (o *Location) UnsetNameAscii()`

UnsetNameAscii ensures that no value is present for NameAscii, not even an explicit nil
### GetLat

`func (o *Location) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Location) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Location) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *Location) HasLat() bool`

HasLat returns a boolean if a field has been set.

### SetLatNil

`func (o *Location) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *Location) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLon

`func (o *Location) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *Location) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *Location) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *Location) HasLon() bool`

HasLon returns a boolean if a field has been set.

### SetLonNil

`func (o *Location) SetLonNil(b bool)`

 SetLonNil sets the value for Lon to be an explicit nil

### UnsetLon
`func (o *Location) UnsetLon()`

UnsetLon ensures that no value is present for Lon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


