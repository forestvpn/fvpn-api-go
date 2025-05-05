# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartDate** | **time.Time** | Date when the subscription was first created.The date might differ from the created date due to backdating. | 
**EndedAt** | Pointer to **NullableTime** | If the subscription has ended, the date the subscription ended. | [optional] 
**CurrentPeriodStart** | **time.Time** | Start of the current period that the subscription has been invoiced for. | 
**CurrentPeriodEnd** | **time.Time** | End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created. | 
**TrialStart** | Pointer to **NullableTime** | If the subscription has a trial, the beginning of that trial. | [optional] 
**TrialEnd** | Pointer to **NullableTime** | If the subscription has a trial, the end of that trial. | [optional] 
**CancelAt** | Pointer to **NullableTime** | A date in the future at which the subscription will automatically get canceled. | [optional] 
**CanceledAt** | Pointer to **NullableTime** | If the subscription has been canceled, the date of that cancellation. | [optional] 
**Status** | Pointer to [**SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] 
**Description** | Pointer to **NullableString** | The subscription&#39;s description, meant to be displayable to the customer.Use this field to optionally store an explanation of the subscription. | [optional] 
**ManagedBy** | [**NullableManagedBy**](ManagedBy.md) |  | 
**Amount** | **float64** |  | [readonly] 
**Currency** | **string** |  | 
**LatestInvoice** | [**NullableInvoice**](Invoice.md) |  | [readonly] 
**PendingUpdate** | [**NullableSubscriptionPendingUpdate**](SubscriptionPendingUpdate.md) |  | [readonly] 
**Plan** | [**Plan**](Plan.md) |  | [readonly] 
**PlanPrice** | [**PlanPrice**](PlanPrice.md) |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewSubscription

`func NewSubscription(id string, startDate time.Time, currentPeriodStart time.Time, currentPeriodEnd time.Time, managedBy NullableManagedBy, amount float64, currency string, latestInvoice NullableInvoice, pendingUpdate NullableSubscriptionPendingUpdate, plan Plan, planPrice PlanPrice, createdAt time.Time, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetStartDate

`func (o *Subscription) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Subscription) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Subscription) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndedAt

`func (o *Subscription) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Subscription) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Subscription) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Subscription) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *Subscription) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Subscription) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetCurrentPeriodStart

`func (o *Subscription) GetCurrentPeriodStart() time.Time`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *Subscription) GetCurrentPeriodStartOk() (*time.Time, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *Subscription) SetCurrentPeriodStart(v time.Time)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### GetCurrentPeriodEnd

`func (o *Subscription) GetCurrentPeriodEnd() time.Time`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *Subscription) GetCurrentPeriodEndOk() (*time.Time, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *Subscription) SetCurrentPeriodEnd(v time.Time)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetTrialStart

`func (o *Subscription) GetTrialStart() time.Time`

GetTrialStart returns the TrialStart field if non-nil, zero value otherwise.

### GetTrialStartOk

`func (o *Subscription) GetTrialStartOk() (*time.Time, bool)`

GetTrialStartOk returns a tuple with the TrialStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStart

`func (o *Subscription) SetTrialStart(v time.Time)`

SetTrialStart sets TrialStart field to given value.

### HasTrialStart

`func (o *Subscription) HasTrialStart() bool`

HasTrialStart returns a boolean if a field has been set.

### SetTrialStartNil

`func (o *Subscription) SetTrialStartNil(b bool)`

 SetTrialStartNil sets the value for TrialStart to be an explicit nil

### UnsetTrialStart
`func (o *Subscription) UnsetTrialStart()`

UnsetTrialStart ensures that no value is present for TrialStart, not even an explicit nil
### GetTrialEnd

`func (o *Subscription) GetTrialEnd() time.Time`

GetTrialEnd returns the TrialEnd field if non-nil, zero value otherwise.

### GetTrialEndOk

`func (o *Subscription) GetTrialEndOk() (*time.Time, bool)`

GetTrialEndOk returns a tuple with the TrialEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialEnd

`func (o *Subscription) SetTrialEnd(v time.Time)`

SetTrialEnd sets TrialEnd field to given value.

### HasTrialEnd

`func (o *Subscription) HasTrialEnd() bool`

HasTrialEnd returns a boolean if a field has been set.

### SetTrialEndNil

`func (o *Subscription) SetTrialEndNil(b bool)`

 SetTrialEndNil sets the value for TrialEnd to be an explicit nil

### UnsetTrialEnd
`func (o *Subscription) UnsetTrialEnd()`

UnsetTrialEnd ensures that no value is present for TrialEnd, not even an explicit nil
### GetCancelAt

`func (o *Subscription) GetCancelAt() time.Time`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *Subscription) GetCancelAtOk() (*time.Time, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *Subscription) SetCancelAt(v time.Time)`

SetCancelAt sets CancelAt field to given value.

### HasCancelAt

`func (o *Subscription) HasCancelAt() bool`

HasCancelAt returns a boolean if a field has been set.

### SetCancelAtNil

`func (o *Subscription) SetCancelAtNil(b bool)`

 SetCancelAtNil sets the value for CancelAt to be an explicit nil

### UnsetCancelAt
`func (o *Subscription) UnsetCancelAt()`

UnsetCancelAt ensures that no value is present for CancelAt, not even an explicit nil
### GetCanceledAt

`func (o *Subscription) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *Subscription) GetCanceledAtOk() (*time.Time, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *Subscription) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### HasCanceledAt

`func (o *Subscription) HasCanceledAt() bool`

HasCanceledAt returns a boolean if a field has been set.

### SetCanceledAtNil

`func (o *Subscription) SetCanceledAtNil(b bool)`

 SetCanceledAtNil sets the value for CanceledAt to be an explicit nil

### UnsetCanceledAt
`func (o *Subscription) UnsetCanceledAt()`

UnsetCanceledAt ensures that no value is present for CanceledAt, not even an explicit nil
### GetStatus

`func (o *Subscription) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *Subscription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Subscription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Subscription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Subscription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Subscription) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Subscription) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetManagedBy

`func (o *Subscription) GetManagedBy() ManagedBy`

GetManagedBy returns the ManagedBy field if non-nil, zero value otherwise.

### GetManagedByOk

`func (o *Subscription) GetManagedByOk() (*ManagedBy, bool)`

GetManagedByOk returns a tuple with the ManagedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedBy

`func (o *Subscription) SetManagedBy(v ManagedBy)`

SetManagedBy sets ManagedBy field to given value.


### SetManagedByNil

`func (o *Subscription) SetManagedByNil(b bool)`

 SetManagedByNil sets the value for ManagedBy to be an explicit nil

### UnsetManagedBy
`func (o *Subscription) UnsetManagedBy()`

UnsetManagedBy ensures that no value is present for ManagedBy, not even an explicit nil
### GetAmount

`func (o *Subscription) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Subscription) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Subscription) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Subscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Subscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Subscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetLatestInvoice

`func (o *Subscription) GetLatestInvoice() Invoice`

GetLatestInvoice returns the LatestInvoice field if non-nil, zero value otherwise.

### GetLatestInvoiceOk

`func (o *Subscription) GetLatestInvoiceOk() (*Invoice, bool)`

GetLatestInvoiceOk returns a tuple with the LatestInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestInvoice

`func (o *Subscription) SetLatestInvoice(v Invoice)`

SetLatestInvoice sets LatestInvoice field to given value.


### SetLatestInvoiceNil

`func (o *Subscription) SetLatestInvoiceNil(b bool)`

 SetLatestInvoiceNil sets the value for LatestInvoice to be an explicit nil

### UnsetLatestInvoice
`func (o *Subscription) UnsetLatestInvoice()`

UnsetLatestInvoice ensures that no value is present for LatestInvoice, not even an explicit nil
### GetPendingUpdate

`func (o *Subscription) GetPendingUpdate() SubscriptionPendingUpdate`

GetPendingUpdate returns the PendingUpdate field if non-nil, zero value otherwise.

### GetPendingUpdateOk

`func (o *Subscription) GetPendingUpdateOk() (*SubscriptionPendingUpdate, bool)`

GetPendingUpdateOk returns a tuple with the PendingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingUpdate

`func (o *Subscription) SetPendingUpdate(v SubscriptionPendingUpdate)`

SetPendingUpdate sets PendingUpdate field to given value.


### SetPendingUpdateNil

`func (o *Subscription) SetPendingUpdateNil(b bool)`

 SetPendingUpdateNil sets the value for PendingUpdate to be an explicit nil

### UnsetPendingUpdate
`func (o *Subscription) UnsetPendingUpdate()`

UnsetPendingUpdate ensures that no value is present for PendingUpdate, not even an explicit nil
### GetPlan

`func (o *Subscription) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Subscription) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Subscription) SetPlan(v Plan)`

SetPlan sets Plan field to given value.


### GetPlanPrice

`func (o *Subscription) GetPlanPrice() PlanPrice`

GetPlanPrice returns the PlanPrice field if non-nil, zero value otherwise.

### GetPlanPriceOk

`func (o *Subscription) GetPlanPriceOk() (*PlanPrice, bool)`

GetPlanPriceOk returns a tuple with the PlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanPrice

`func (o *Subscription) SetPlanPrice(v PlanPrice)`

SetPlanPrice sets PlanPrice field to given value.


### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


