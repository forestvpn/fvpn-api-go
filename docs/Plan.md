# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**IsMostPopular** | Pointer to **bool** |  | [optional] 
**Prices** | [**[]PlanPrice**](PlanPrice.md) |  | 
**Currencies** | **NullableString** |  | [readonly] 

## Methods

### NewPlan

`func NewPlan(id string, slug string, name string, prices []PlanPrice, currencies NullableString, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *Plan) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Plan) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Plan) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Plan) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Plan) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIsMostPopular

`func (o *Plan) GetIsMostPopular() bool`

GetIsMostPopular returns the IsMostPopular field if non-nil, zero value otherwise.

### GetIsMostPopularOk

`func (o *Plan) GetIsMostPopularOk() (*bool, bool)`

GetIsMostPopularOk returns a tuple with the IsMostPopular field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMostPopular

`func (o *Plan) SetIsMostPopular(v bool)`

SetIsMostPopular sets IsMostPopular field to given value.

### HasIsMostPopular

`func (o *Plan) HasIsMostPopular() bool`

HasIsMostPopular returns a boolean if a field has been set.

### GetPrices

`func (o *Plan) GetPrices() []PlanPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *Plan) GetPricesOk() (*[]PlanPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *Plan) SetPrices(v []PlanPrice)`

SetPrices sets Prices field to given value.


### SetPricesNil

`func (o *Plan) SetPricesNil(b bool)`

 SetPricesNil sets the value for Prices to be an explicit nil

### UnsetPrices
`func (o *Plan) UnsetPrices()`

UnsetPrices ensures that no value is present for Prices, not even an explicit nil
### GetCurrencies

`func (o *Plan) GetCurrencies() string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *Plan) GetCurrenciesOk() (*string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *Plan) SetCurrencies(v string)`

SetCurrencies sets Currencies field to given value.


### SetCurrenciesNil

`func (o *Plan) SetCurrenciesNil(b bool)`

 SetCurrenciesNil sets the value for Currencies to be an explicit nil

### UnsetCurrencies
`func (o *Plan) UnsetCurrencies()`

UnsetCurrencies ensures that no value is present for Currencies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


