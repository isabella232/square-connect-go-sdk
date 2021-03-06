# Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for the payment. | [optional] [default to null]
**CreatedAt** | **string** | Timestamp of when the payment was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | Timestamp of when the payment was last updated, in RFC 3339 format. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TipMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ProcessingFee** | [**[]ProcessingFee**](ProcessingFee.md) | Processing fees and fee adjustments assessed by Square on this payment. | [optional] [default to null]
**RefundedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Status** | **string** | Indicates whether the payment is &#x60;APPROVED&#x60;, &#x60;COMPLETED&#x60;, &#x60;CANCELED&#x60;, or &#x60;FAILED&#x60;. | [optional] [default to null]
**DelayDuration** | **string** | The duration of time after the payment&#x27;s creation when Square automatically applies the &#x60;delay_action&#x60; to the payment. This automatic &#x60;delay_action&#x60; applies only to payments that don&#x27;t reach a terminal state (COMPLETED, CANCELED, or FAILED) before the &#x60;delay_duration&#x60; time period.  This field is specified as a time duration, in RFC 3339 format.  Notes: This feature is only supported for card payments.  Default:  - Card Present payments: \&quot;PT36H\&quot; (36 hours) from the creation time. - Card Not Present payments: \&quot;P7D\&quot; (7 days) from the creation time. | [optional] [default to null]
**DelayAction** | **string** | The action to be applied to the payment when the &#x60;delay_duration&#x60; has elapsed. This field is read only.  Current values include: &#x60;CANCEL&#x60; | [optional] [default to null]
**DelayedUntil** | **string** | Read only timestamp of when the &#x60;delay_action&#x60; will automatically be applied, in RFC 3339 format.  Note that this field is calculated by summing the payment&#x27;s &#x60;delay_duration&#x60; and &#x60;created_at&#x60; fields. The &#x60;created_at&#x60; field is generated by Square and may not exactly match the time on your local machine. | [optional] [default to null]
**SourceType** | **string** | The source type for this payment  Current values include: &#x60;CARD&#x60;. | [optional] [default to null]
**CardDetails** | [***CardPaymentDetails**](CardPaymentDetails.md) |  | [optional] [default to null]
**LocationId** | **string** | ID of the location associated with the payment. | [optional] [default to null]
**OrderId** | **string** | ID of the order associated with this payment. | [optional] [default to null]
**ReferenceId** | **string** | An optional ID that associates this payment with an entity in another system. | [optional] [default to null]
**CustomerId** | **string** | The [Customer](#type-customer) ID of the customer associated with the payment. | [optional] [default to null]
**EmployeeId** | **string** | An optional ID of the employee associated with taking this payment. | [optional] [default to null]
**RefundIds** | **[]string** | List of &#x60;refund_id&#x60;s identifying refunds for this payment. | [optional] [default to null]
**BuyerEmailAddress** | **string** | The buyer&#x27;s e-mail address | [optional] [default to null]
**BillingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Note** | **string** | An optional note to include when creating a payment | [optional] [default to null]
**StatementDescriptionIdentifier** | **string** | Additional payment information that gets added on the customer&#x27;s card statement as part of the statement description.  Note that the &#x60;statement_description_identifier&#x60; may get truncated on the statement description to fit the required information including the Square identifier (SQ *) and name of the merchant taking the payment. | [optional] [default to null]
**ReceiptNumber** | **string** | The payment&#x27;s receipt number. The field will be missing if a payment is CANCELED | [optional] [default to null]
**ReceiptUrl** | **string** | The URL for the payment&#x27;s receipt. The field will only be populated for COMPLETED payments. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

