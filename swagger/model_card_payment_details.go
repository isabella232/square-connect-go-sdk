/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Reflects the current status of a card payment.
type CardPaymentDetails struct {
	// The card payment's current state. It can be one of: `AUTHORIZED`, `CAPTURED`, `VOIDED`, `FAILED`.
	Status string `json:"status,omitempty"`
	Card   *Card  `json:"card,omitempty"`
	// The method used to enter the card's details for the payment.  Can be `KEYED`, `SWIPED`, `EMV`, `ON_FILE`, or `CONTACTLESS`.
	EntryMethod string `json:"entry_method,omitempty"`
	// Status code returned from the Card Verification Value (CVV) check. Can be `CVV_ACCEPTED`, `CVV_REJECTED`, `CVV_NOT_CHECKED`.
	CvvStatus string `json:"cvv_status,omitempty"`
	// Status code returned from the Address Verification System (AVS) check. Can be `AVS_ACCEPTED`, `AVS_REJECTED`, `AVS_NOT_CHECKED`.
	AvsStatus string `json:"avs_status,omitempty"`
	// Status code returned by the card issuer that describes the payment's authorization status.
	AuthResultCode string `json:"auth_result_code,omitempty"`
	// For EMV payments, identifies the EMV application used for the payment.
	ApplicationIdentifier string `json:"application_identifier,omitempty"`
	// For EMV payments, the human-readable name of the EMV application used for the payment.
	ApplicationName string `json:"application_name,omitempty"`
	// For EMV payments, the cryptogram generated for the payment.
	ApplicationCryptogram string `json:"application_cryptogram,omitempty"`
	// For EMV payments, method used to verify the cardholder's identity.  Can be one of `PIN`, `SIGNATURE`, `PIN_AND_SIGNATURE`, `ON_DEVICE`, or `NONE`.
	VerificationMethod string `json:"verification_method,omitempty"`
	// For EMV payments, the results of the cardholder verification.  Can be one of `SUCCESS`, `FAILURE`, or `UNKNOWN`.
	VerificationResults string `json:"verification_results,omitempty"`
	// The statement description sent to the card networks.  Note: The actual statement description will vary and is likely to be truncated and appended with additional information on a per issuer basis.
	StatementDescription string         `json:"statement_description,omitempty"`
	DeviceDetails        *DeviceDetails `json:"device_details,omitempty"`
	// Whether or not the card is required to be physically present in order for the payment to be refunded.  If true, the card is required to be present.
	RefundRequiresCardPresence bool `json:"refund_requires_card_presence,omitempty"`
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
}