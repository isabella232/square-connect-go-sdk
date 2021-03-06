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

// Represents the details of a tender with `type` `BANK_TRANSFER`.  See [PaymentBankTransferDetails](#type-paymentbanktransferdetails) for more exposed details of a bank transfer payment.
type TenderBankTransferDetails struct {
	Status *TenderBankTransferDetailsStatus `json:"status,omitempty"`
}
