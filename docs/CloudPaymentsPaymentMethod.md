# CloudPaymentsPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to **NullableString** |  | [optional] 
**Card** | [**CloudPaymentsCardDetails**](CloudPaymentsCardDetails.md) |  | 
**IsExpired** | **bool** |  | [readonly] 
**NeedAuthentication** | **bool** |  | [readonly] 
**ThreeDs** | [**CloudPaymentsThreeDSecure**](CloudPaymentsThreeDSecure.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewCloudPaymentsPaymentMethod

`func NewCloudPaymentsPaymentMethod(id string, card CloudPaymentsCardDetails, isExpired bool, needAuthentication bool, threeDs CloudPaymentsThreeDSecure, createdAt time.Time, ) *CloudPaymentsPaymentMethod`

NewCloudPaymentsPaymentMethod instantiates a new CloudPaymentsPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudPaymentsPaymentMethodWithDefaults

`func NewCloudPaymentsPaymentMethodWithDefaults() *CloudPaymentsPaymentMethod`

NewCloudPaymentsPaymentMethodWithDefaults instantiates a new CloudPaymentsPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudPaymentsPaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudPaymentsPaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudPaymentsPaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CloudPaymentsPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudPaymentsPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudPaymentsPaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudPaymentsPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CloudPaymentsPaymentMethod) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CloudPaymentsPaymentMethod) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCard

`func (o *CloudPaymentsPaymentMethod) GetCard() CloudPaymentsCardDetails`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *CloudPaymentsPaymentMethod) GetCardOk() (*CloudPaymentsCardDetails, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *CloudPaymentsPaymentMethod) SetCard(v CloudPaymentsCardDetails)`

SetCard sets Card field to given value.


### GetIsExpired

`func (o *CloudPaymentsPaymentMethod) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *CloudPaymentsPaymentMethod) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *CloudPaymentsPaymentMethod) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.


### GetNeedAuthentication

`func (o *CloudPaymentsPaymentMethod) GetNeedAuthentication() bool`

GetNeedAuthentication returns the NeedAuthentication field if non-nil, zero value otherwise.

### GetNeedAuthenticationOk

`func (o *CloudPaymentsPaymentMethod) GetNeedAuthenticationOk() (*bool, bool)`

GetNeedAuthenticationOk returns a tuple with the NeedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedAuthentication

`func (o *CloudPaymentsPaymentMethod) SetNeedAuthentication(v bool)`

SetNeedAuthentication sets NeedAuthentication field to given value.


### GetThreeDs

`func (o *CloudPaymentsPaymentMethod) GetThreeDs() CloudPaymentsThreeDSecure`

GetThreeDs returns the ThreeDs field if non-nil, zero value otherwise.

### GetThreeDsOk

`func (o *CloudPaymentsPaymentMethod) GetThreeDsOk() (*CloudPaymentsThreeDSecure, bool)`

GetThreeDsOk returns a tuple with the ThreeDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeDs

`func (o *CloudPaymentsPaymentMethod) SetThreeDs(v CloudPaymentsThreeDSecure)`

SetThreeDs sets ThreeDs field to given value.


### GetCreatedAt

`func (o *CloudPaymentsPaymentMethod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudPaymentsPaymentMethod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudPaymentsPaymentMethod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


