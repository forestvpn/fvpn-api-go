# AccountTokenPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The 16-digit account number of the user, with optional separators. | 

## Methods

### NewAccountTokenPairRequest

`func NewAccountTokenPairRequest(number string, ) *AccountTokenPairRequest`

NewAccountTokenPairRequest instantiates a new AccountTokenPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountTokenPairRequestWithDefaults

`func NewAccountTokenPairRequestWithDefaults() *AccountTokenPairRequest`

NewAccountTokenPairRequestWithDefaults instantiates a new AccountTokenPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *AccountTokenPairRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *AccountTokenPairRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *AccountTokenPairRequest) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


